package main

import (
	"github.com/lzxm160/csvtomysql/utils"
	"flag"
	"fmt"
	"os"
)
type config struct {
	MysqlConnectString string`json:"mysqlConnectString,omitempty"`
}

var(
	configPath string
	csvPath string
)

func init() {
	flag.Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr,
			"usage: server -config=[string]\n")
		flag.PrintDefaults()
		os.Exit(2)
	}
	flag.StringVar(&configPath, "config", "config.toml", "path of config file")
	flag.StringVar(&csvPath, "csv", "csv.csv", "path of csv file")
	flag.Parse()
}
func main() {
	var cfg config
	utils.LoadConfig(configPath,&cfg)
	fmt.Println("::",cfg)

	//select {}
}
