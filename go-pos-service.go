// Package main provides ...
package main

import (
	"fmt"
	"time"

	"github.com/duminghui/go-pos-service/dstcmd"
	"github.com/duminghui/go-util/utime"
)

func main() {
	// tx, _ := dstcmd.GetRawTransaction("d1a22903ea81ffaa0dfd1ef7e3878216c662a6587e14f06ef6ae49d0a2ba5f50")
	// tx, _ := dstcmd.GetRawTransaction("01f44544946d810d5c2bcb716edd4ed9dae7caba36c2104d37b03308c008f315")
	// vout := tx.Vout
	// tx, _ := dstcmd.GetRawTransaction("826b590f997c4366c51a4e1fdf0ad272b1c9b27886e289cc3c341807af9f26ba")
	// log.Printf("Size of TX: %d", unsafe.Sizeof(*tx))
	// fmt.Printf("%s\n", tx)
	// tx := dstcmd.GetRawTransaction("iTTTT")
	// fmt.Println(tx == nil)
	// for i, vout := range *tx.Vout {
	// fmt.Printf("%d:%s\n", i, vout)
	// }
	// fmt.Printf("%#v\n", config.DBConfig)
	// user := dbmodel.FindUserByID("385061500034875392")
	// user := dbmodel.FindUserByID("11111")
	// fmt.Printf("%#v\n", user)
	// fmt.Println(user == nil)
	// fmt.Println(user.ID)
	// userAddr := dbmodel.FindUserByAddr("fMZictACJc9dKhrMxCKRMkWNpq8Ni2ZBx8")
	// userAddr := dbmodel.FindUserByAddr("AAA")
	// fmt.Printf("%#v\n", userAddr)
	// fmt.Println(userAddr == nil)
	// latestTx := dbmodel.FindLatestTransaction()
	// fmt.Printf("%#v\n", latestTx)
	// fmt.Println(latestTx == nil)
	// tx := new(dbmodel.DstTransaction)
	// tx.Txid = "aa"
	// tx.Category = "bbb"
	// tx.Amount = 10.10001
	// tx.Txtime = 0
	// tx.Save()
	// tx.Delete()
	// txs, _ := dstcmd.ListTransactions(0, 1)
	// fmt.Printf("%p\n", txs)
	// tmp = append(tmp, 6)
	// fmt.Println(tmp)
	// fmt.Printf("%p\n", &tmp)
	startTime := time.Now()
	// index, count := dstcmd.TxIndexCountByID("d407fcca75f4eea5e0466f6d72e3ba2d0b1202d4400da0e72a43ea9cbbc4655c")
	// index, count := dstcmd.TxIndexCountByID("58f45c68e156ffe94e17c7c5d57c89e7b872445b9c23c4ea18e3fa2a32c1567c")
	// index, count := dstcmd.TxIndexCountByID("d1a22903ea81ffaa0dfd1ef7e3878216c662a6587e14f06ef6ae49d0a2ba5f50")
	// fmt.Println(index, count)
	// fmt.Printf("%s\n", time.Since(startTime))
	// startTime = time.Now()
	// index, count = dstcmd.TxIndexCountByIDSync("d407fcca75f4eea5e0466f6d72e3ba2d0b1202d4400da0e72a43ea9cbbc4655c")
	// index, count = dstcmd.TxIndexCountByIDSync("58f45c68e156ffe94e17c7c5d57c89e7b872445b9c23c4ea18e3fa2a32c1567c")
	// index, count := dstcmd.TxIndexCountByIDSync("d1a22903ea81ffaa0dfd1ef7e3878216c662a6587e14f06ef6ae49d0a2ba5f50")
	index, count := dstcmd.TxIndexCountByIDSync("972f1e6ec7c025fc524bf224192c9aa0b8f87e10b9c11a277466097a5ad7ec92")
	// index, count := dstcmd.TxIndexCountByID("972f1e6ec7c025fc524bf224192c9aa0b8f87e10b9c11a277466097a5ad7ec92")
	fmt.Println(index, count)
	fmt.Printf("%s\n", time.Since(startTime))
	fmt.Println(utime.FormatLongTimeStrUTC(100000))
}
