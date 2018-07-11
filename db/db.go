// Package db provides ...
package db

import (
	"fmt"
	"log"

	"github.com/duminghui/go-pos-service/config"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func test() {}
func init() {
	dbConfig := config.DBConfig
	mysqlConfig := mysql.NewConfig()
	mysqlConfig.Net = dbConfig.Net
	mysqlConfig.Addr = dbConfig.Addr
	mysqlConfig.DBName = dbConfig.DBName
	mysqlConfig.User = dbConfig.User
	mysqlConfig.Passwd = dbConfig.Passwd
	mysqlDSN := mysqlConfig.FormatDSN()
	log.Printf("Mysql DSN: %s\n", mysqlDSN)
	db, err := sqlx.Connect("mysql", mysqlDSN)
	if err != nil {
		panic(fmt.Sprintf("Open Mysql Error:%s\n", err))
	}
	DB = db
	log.Printf("Mysql Status:%v\n", db.Stats())
}
