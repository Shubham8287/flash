package conf

import (
	"encoding/json"
	. "flash/logger"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
)

const CONFIGPATH = "./conf.json"

type Config struct {
	Algo     string `json:"algo"`
	DataPath string `json:"data_path"`
}

var configInstance Config

func readConfig() Config {
	Log.Info("config_reading_start")
	jsonFile, err := os.Open(CONFIGPATH)
	if err != nil {
		fmt.Println("error:", err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()

	config := Config{}
	json.Unmarshal(byteValue, &config)
	Log.WithFields(logrus.Fields{
		"algorithm": config.Algo,
		"data_path": config.DataPath,
	}).Info("config_reading_done")

	return config
}

func Get() Config {
	if configInstance == (Config{}) {
		configInstance = readConfig()
	}
	return configInstance
}
