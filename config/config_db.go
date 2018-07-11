// Package config provides ...
package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

const (
	db_config_file = "./config-db.json"
)

type dbConfig struct {
	Net    string `json:"net"`
	Addr   string `json:"addr"`
	DBName string `json:"dbname"`
	User   string `json:"user"`
	Passwd string `json:"passwd"`
}

var DBConfig *dbConfig

func init() {
	configBytes, err := ioutil.ReadFile(db_config_file)
	if err != nil {
		log.Fatal(err)
	}
	DBConfig = &dbConfig{}
	err = json.Unmarshal(configBytes, DBConfig)
	if err != nil {
		log.Fatal(err)
	}
}
