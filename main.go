package main

import (
	"github.com/lzxm160/csvtomysql/utils"
	"flag"
	"fmt"
)
type config struct {
	mysqlConnectString string
}
func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "config.toml", "path of config file")
	flag.Parse()
	var cfg config
	utils.LoadConfig(configPath,&cfg)
	fmt.Println("::",cfg)

	select {}
}
