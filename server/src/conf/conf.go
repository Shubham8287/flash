package conf

import (
    "encoding/json"
    "fmt"
	"os"
	"io/ioutil"
)

const CONFIGPATH = "conf/conf.json";

type Config struct {
    Algo string `json:"algo"`
}

func readConfig() Config {
	jsonFile, err := os.Open(CONFIGPATH)
	if err != nil {
		fmt.Println("error:", err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()

	config := Config{}
	json.Unmarshal(byteValue, &config)

	return config;
}

func Init() {
	config := readConfig();

	if(config.Algo == "map") {

	} else {
		
	}
	fmt.Println(config.Algo)
}