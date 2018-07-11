// Package dbmodel provides ...
package dbmodel

import (
	"database/sql"
	"fmt"
	"log"
)

func checkQueryErr(err error, sqlStr, noRowMsg string) bool {
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(noRowMsg)
			return true
		}
		panic(fmt.Sprintf("%s[%s]", err, sqlStr))
	}
	return false
}

func checkIUDErr(err error, sqlStr string) {
	if err != nil {
		panic(fmt.Sprintf("%s[%s]", err, sqlStr))
	}
}
