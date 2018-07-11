// Package dstcmd provides ...
package dstcmd

import (
	"fmt"
	"sort"
	"sync"

	"github.com/francoispqt/gojay"
)

const (
	TX_CATEGORY_RECEIVE    = "receive"
	TX_CATEGORY_SEND       = "send"
	TX_CATEGORY_SENDTOSELF = "sendtoself"
	TX_CATEGORY_IMMATURE   = "immature"
	TX_CATEGORY_GENERATE   = "generate"
)

type listTransaction struct {
	Category      string  `json:"category"`
	Amount        float64 `json:"amount"`
	Confirmations int     `json:"confirmations"`
	Generated     bool    `json:"generated"`
	Blocktime     int     `json:"blocktime"`
	Txid          string  `json:"txid"`
}

func (l *listTransaction) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "category":
		return dec.String(&l.Category)
	case "amount":
		return dec.Float(&l.Amount)
	case "confirmations":
		return dec.Int(&l.Confirmations)
	case "generated":
		return dec.Bool(&l.Generated)
	case "blocktime":
		return dec.Int(&l.Blocktime)
	case "txid":
		return dec.String(&l.Txid)
	}
	return nil
}

func (l *listTransaction) NKeys() int {
	return 0
}

type listTransactionArray []*listTransaction

func (l *listTransactionArray) UnmarshalJSONArray(dec *gojay.Decoder) error {
	transaction := new(listTransaction)
	if err := dec.Object(transaction); err != nil {
		return err
	}
	*l = append(*l, transaction)
	return nil
}

func ListTransactions(from, count int) (listTransactionArray, error) {
	cmd := fmt.Sprintf("listtransactions airdrop %v %v", count, from)
	bytes, _ := execDstShell(cmd)
	// listTXArray := listTransactionArray{}
	listTXArray := make(listTransactionArray, 0, count)
	err := gojay.UnmarshalJSONArray(*bytes, &listTXArray)
	return listTXArray, err
}

func TxIndexCountByID(txid string) (index, count int) {
	txs, _ := ListTransactions(0, 1)
	if len(txs) == 0 {
		return -1, 0
	}
	const STEP = 300
	index, count, find := -1, 0, false
	find = true
	startTxID := txs[0].Txid
	for {
		txs, _ = ListTransactions(count, STEP)
		countTmp := len(txs)
		count += countTmp
		if countTmp == 0 {
			break
		}
		if !find {
			for i := countTmp - 1; i >= 0; i-- {
				index++
				if txs[i].Txid == txid {
					find = true
					break
				}
			}
		}
	}
	txs, _ = ListTransactions(0, 1)
	if startTxID != txs[0].Txid {
		return TxIndexCountByID(txid)
	}
	return index, count
}

type indexQueue struct {
	queue chan int
	done  chan struct{}
}

func newIndexQueue() *indexQueue {
	ia := &indexQueue{
		queue: make(chan int),
		done:  make(chan struct{}),
	}
	ia.process()
	return ia
}

func (ia *indexQueue) process() {
	index := 0
	go func() {
		for {
			select {
			case ia.queue <- index:
				index++
			case <-ia.done:
				return
			}
		}
	}()
}

func (ia *indexQueue) get() int {
	return <-ia.queue
}

func (ia *indexQueue) close() {
	ia.done <- struct{}{}
	close(ia.queue)
	close(ia.done)
}

type isFoundIndex struct {
	set  chan struct{}
	get  chan bool
	done chan struct{}
}

func newIsFoundIndex() *isFoundIndex {
	isFoundIndex := &isFoundIndex{
		set:  make(chan struct{}),
		get:  make(chan bool),
		done: make(chan struct{}),
	}
	isFoundIndex.process()
	return isFoundIndex
}

func (ifi *isFoundIndex) setTrue() {
	ifi.set <- struct{}{}
}

func (ifi *isFoundIndex) is() bool {
	return <-ifi.get
}

func (ifi *isFoundIndex) process() {
	isFoundIndex := false
	go func() {
		for {
			select {
			case <-ifi.set:
				isFoundIndex = true
			case ifi.get <- isFoundIndex:
			case <-ifi.done:
				return
			}
		}
	}()
}

func (ifi *isFoundIndex) close() {
	ifi.done <- struct{}{}
	close(ifi.set)
	close(ifi.get)
	close(ifi.done)
}

type reqQueue chan struct{}

type reqQueQueue chan reqQueue

type indexCount struct {
	index int
	count int
}

type respQueue chan indexCount

type processData struct {
	txid         string
	indexQueue   *indexQueue
	isFoundIndex *isFoundIndex
	reqQueue     reqQueue
	reqQueQueue  reqQueQueue
	respQueue    respQueue
	n            *sync.WaitGroup
}

func txIndexCountByIDSub(processData *processData) {
	defer func() {
		close(processData.reqQueue)
		processData.n.Done()
	}()
	const STEP = 300
	for {
		<-processData.reqQueue
		index := processData.indexQueue.get()
		start := index * STEP
		txs, _ := ListTransactions(start, STEP)
		txsLen := len(txs)
		if txsLen == 0 {
			return
		}
		txIndex := -1
		if !processData.isFoundIndex.is() {
			for i, j := txsLen-1, 0; i >= 0; i, j = i-1, j+1 {
				if processData.txid == txs[i].Txid {
					txIndex = start + j
					// log.Println(txIndex)
					break
				}
			}
		}
		processData.respQueue <- indexCount{txIndex, txsLen}
		if txsLen < STEP {
			return
		}
		processData.reqQueQueue <- processData.reqQueue
		<-processData.reqQueQueue <- struct{}{}
	}
}

type indexs []int

func (is indexs) Len() int {
	return len(is)
}

func (is indexs) Less(i int, j int) bool {
	return is[i] < is[j]
}

func (is indexs) Swap(i int, j int) {
	is[i], is[j] = is[j], is[i]
}

func TxIndexCountByIDSync(txid string) (int, int) {
	txs, _ := ListTransactions(0, 1)
	if len(txs) == 0 {
		return -1, 0
	}
	startTxID := txs[0].Txid
	indexQueue := newIndexQueue()
	isFoundIndex := newIsFoundIndex()
	respQueue := make(respQueue)
	const chanNums = 3
	reqQueQueue := make(reqQueQueue, chanNums)
	var n sync.WaitGroup
	for i := 0; i < chanNums; i++ {
		processData := new(processData)
		processData.txid = txid
		reqQueue := make(reqQueue, 1)
		reqQueQueue <- reqQueue
		processData.reqQueue = reqQueue
		processData.reqQueQueue = reqQueQueue
		processData.respQueue = respQueue
		processData.indexQueue = indexQueue
		processData.isFoundIndex = isFoundIndex
		processData.n = &n
		n.Add(1)
		go txIndexCountByIDSub(processData)
	}
	resultResp := make(chan indexCount)
	go func() {
		is := indexs{}
		count := 0
		for resp := range respQueue {
			if resp.index != -1 {
				is = append(is, resp.index)
			}
			count += resp.count
		}
		index := -1
		if len(is) > 0 {
			if len(is) > 1 {
				sort.Sort(is)
			}
			index = is[0]
		}
		resultResp <- indexCount{index, count}
	}()
	for i := 0; i < chanNums; i++ {
		<-reqQueQueue <- struct{}{}
	}
	n.Wait()
	indexQueue.close()
	// isFoundIndex.close()
	close(respQueue)
	close(reqQueQueue)
	resp := <-resultResp
	close(resultResp)
	txs, _ = ListTransactions(0, 1)
	if startTxID != txs[0].Txid {
		return TxIndexCountByID(txid)
	}
	return resp.index, resp.count
}
