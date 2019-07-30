package main

import (
	"database/sql"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

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
	logPath      string
	line         uint64
	sleepSeconds uint64
	l            *log.Logger
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
	flag.StringVar(&logPath, "log", "log.log", "path of log file")
	flag.Uint64Var(&line, "line", 500000, "line")
	flag.Uint64Var(&sleepSeconds, "sleep", 3, "sleep")
	flag.Parse()

	file, err := os.Create(logPath)
	if err != nil {
		log.Println(err)
		return
	}
	l = log.New(file, " ", log.Lshortfile)
	l.SetOutput(os.Stdout)
}
func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.MysqlConnectString+cfg.DbName+"?autocommit=false")
	if err != nil {
		return nil, err
	}
	return db, nil
}
func readAndWrite(s *sql.DB, cfg config) {
	f, err := os.Open(csvPath)
	if err != nil {
		l.Println(err)
		return
	}
	r := csv.NewReader(f)
	i := uint64(0)
	for {
		records, err := r.Read()
		if err == io.EOF {
			log.Println(err)
			break
		}
		if err != nil {
			l.Println(err)
			return
		}
		if i == 0 {
			i++
			continue
		}
		i++
		if uint64(i)%line == 0 {
			log.Println("sleep for a while")
			time.Sleep(time.Second * time.Duration(sleepSeconds))
		}
		name := records[0]
		email := records[1]
		l.Println(records)
		insertQuery := fmt.Sprintf("INSERT INTO %s (name, email) VALUES ('%s',' %s')", cfg.TableName, name, email)
		if _, err = s.Exec(insertQuery); err != nil {
			l.Println(err)
		}
	}
}
func main() {
	var cfg config
	utils.LoadConfig(configPath, &cfg)
	l.Println("config:", cfg)

	db, err := openDB(cfg)
	defer db.Close()
	if err != nil {
		l.Println(err)
		return
	}

	readAndWrite(db, cfg)
}
