package main

import (
	"database/sql"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lzxm160/csvtomysql/utils"
)

type config struct {
	MysqlConnectString string `json:"mysqlConnectString,omitempty"`
	DbName             string `json:"dbName,omitempty"`
	TableName          string `json:"tableName,omitempty"`
}

var (
	configPath   string
	csvPath      string
	line         uint64
	sleepSeconds uint64
)

func init() {
	flag.Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr,
			"usage: csvtomysql -config=[string] -csv=[string] -line=500000 -sleep=3\n")
		flag.PrintDefaults()
		os.Exit(2)
	}
	flag.StringVar(&configPath, "config", "config.toml", "path of config file")
	flag.StringVar(&csvPath, "csv", "csv.csv", "path of csv file")
	flag.Uint64Var(&line, "line", 500000, "line")
	flag.Uint64Var(&sleepSeconds, "sleep", 3, "sleep")
	flag.Parse()
}
func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.MysqlConnectString+cfg.DbName+"?autocommit=false")
	if err != nil {
		return nil, err
	}
	return db, nil
}
func readAndWrite(s *sql.DB) {
	f, err := os.Open(csvPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
			return
		}

		for value := range record {
			fmt.Printf(" value:%v %v\n", value, record[value])
		}
	}
}
func main() {
	var cfg config
	utils.LoadConfig(configPath, &cfg)
	fmt.Println("config:", cfg)
	db, err := openDB(cfg)
	defer db.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
	readAndWrite(db)
}
