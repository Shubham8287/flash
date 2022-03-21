package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const CONFIGPATH = "./conf.json"

type Config struct {
	Algo     string `json:"algo"`
	DataPath string `json:"data_path"`
}

var configInstance Config

func readConfig() Config {
	jsonFile, err := os.Open(CONFIGPATH)
	if err != nil {
		fmt.Println("error:", err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()

	config := Config{}
	json.Unmarshal(byteValue, &config)

	return config
}

func Get() Config {
	if configInstance == (Config{}) {
		configInstance = readConfig()
	}
	return configInstance
}
