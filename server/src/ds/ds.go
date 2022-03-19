package ds

import (
	"encoding/json"
	"flash/conf"
	"flash/ds/hashmap"
	"flash/ds/trie"
	"fmt"
	"io/ioutil"
)

type DataStructure interface {
	Find(string) []string
	Insert(string)
}

var bucketData DataStructure = nil
var configuration conf.Config = conf.Get()

func init() {
	if configuration.Algo == "map" {
		bucketData = hashmap.Get()
	} else {
		bucketData = trie.Get()
	}
	fillData(bucketData)
}

func fillData(bucketData DataStructure) {
	fmt.Println("reading data...")
	fileData, err := ioutil.ReadFile(configuration.DataPath)
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
		bucketData.Insert(intialWordSet[i])
	}
}

func GetDs() DataStructure {
	if bucketData != nil {
		return bucketData
	}
	return bucketData
}
