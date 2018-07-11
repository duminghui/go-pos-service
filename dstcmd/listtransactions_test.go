package dstcmd

import (
	"fmt"
	"log"
	"testing"
)

func TestListTransactions(t *testing.T) {
	t.SkipNow()
	listTxArray, _ := ListTransactions(0, 10)
	for _, tx := range listTxArray {
		fmt.Printf("TX:%#v\n", tx)
	}
}

func BenchmarkListTransactions(b *testing.B) {
	b.SkipNow()
	for i := 0; i < b.N; i++ {
		ListTransactions(0, 100)
	}
}

func BenchmarkTxIndexCountByTxID(b *testing.B) {
	b.SkipNow()
	for i := 0; i < b.N; i++ {
		TxIndexCountByID("d407fcca75f4eea5e0466f6d72e3ba2d0b1202d4400da0e72a43ea9cbbc4655c")
		// log.Println(index, count)
	}
}

func BenchmarkTxIndexCountByTxID2(b *testing.B) {
	// b.SkipNow()
	for i := 0; i < b.N; i++ {
		index, count := TxIndexCountByIDSync("d407fcca75f4eea5e0466f6d72e3ba2d0b1202d4400da0e72a43ea9cbbc4655c")
		log.Println(index, count)
	}
}
