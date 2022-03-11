package ds

import (
	"encoding/json"
	"flash/conf"
	"flash/ds/hashmap"
	"flash/ds/trie"
	"fmt"
	"io/ioutil"
)

const DataPath = "./data/words_dictionary.json"

type DataStructure interface {
	Find(string) []string
	Insert(string)
}

var DSObj DataStructure = nil

func fillData(DSObj DataStructure) {
	fmt.Println("reading data...")
	fileData, err := ioutil.ReadFile(DataPath)
	if err != nil {
		fmt.Print(err)
	}

	var intialWordSet = []string{}
	fmt.Println("Unmarshalling data...")
	err = json.Unmarshal(fileData, &intialWordSet)
	if err != nil {
		fmt.Println("error:", err)
	}

	for i := 0; i < len(intialWordSet); i++ {
		DSObj.Insert(intialWordSet[i])
	}
}

func GetDs() DataStructure {
	if DSObj != nil {
		return DSObj
	}

	configuration := conf.Get()
	if configuration.Algo == "map" {
		DSObj = hashmap.Get()
	} else {
		DSObj = trie.Get()
	}
	fillData(DSObj)
	return DSObj
}
