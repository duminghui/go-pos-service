// Package dbmodel provides ...
package dbmodel

import (
	"fmt"
	"log"

	"github.com/duminghui/go-pos-service/db"
	"github.com/duminghui/go-util/utime"
)

type DstTransaction struct {
	Txid      string  `db:"txid"`
	Category  string  `db:"category"`
	Amount    float64 `db:"amount"`
	Txtime    int     `db:"txtime"`
	TxtimeStr string  `db:"txtime_str"`
}

func FindLatestTransaction() *DstTransaction {
	const sqlStr = "select * from dst_transaction order by txtime desc limit 1"
	tx := &DstTransaction{}
	err := db.DB.Get(tx, sqlStr)
	if checkQueryErr(err, sqlStr, fmt.Sprintf("No Latest Tx:[%s]", sqlStr)) {
		return nil
	}
	return tx
}

func (tx *DstTransaction) Save() {
	const sqlStr = "insert into dst_transaction(txid,category,amount,txtime,txtime_str) values (:txid,:category,:amount,:txtime,:txtime_str)"
	tx.TxtimeStr = utime.FormatLongTimeStrUTC(tx.Txtime)
	result, err := db.DB.NamedExec(sqlStr, tx)
	checkIUDErr(err, sqlStr)
	rows, _ := result.RowsAffected()
	log.Printf("Insert %d rows to dst_transaction:[%s]", rows, tx.Txid)
}

func (tx *DstTransaction) Delete() {
	DeleteDstTransactionByID(tx.Txid)
}

func DeleteDstTransactionByID(txid string) {
	const sqlStr = "delete from dst_transaction where txid=?"
	result, err := db.DB.Exec(sqlStr, txid)
	checkIUDErr(err, sqlStr)
	rows, _ := result.RowsAffected()
	log.Printf("Delete %d rows in dst_transaction:[%s]", rows, txid)
}
