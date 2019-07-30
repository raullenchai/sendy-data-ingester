package utils

import (
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

func LoadConfig(configFileName string, cfg interface{}) bool {
	var err error
	configFileName, err = filepath.Abs(configFileName)
	if err != nil {
		log.Fatal("LoadConfig", "filepath.Abs err", err)
		return false
	}

	configFile, err := os.Open(configFileName)
	if err != nil {
		log.Fatal("LoadConfig", "os.Open File error: ", err.Error())
		return false
	}
	defer configFile.Close()

	if _, err = toml.DecodeReader(configFile, cfg); err != nil {
		log.Fatal("LoadConfig", "toml.DecodeReader error: ", err.Error())
		return false
	}

	return true
}
