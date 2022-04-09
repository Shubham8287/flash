package ds

import (
	"context"
	"encoding/json"
	"flash/conf"
	"flash/ds/hashmap"
	"flash/ds/trie"
	. "flash/logger"
	"fmt"
	"io/ioutil"
	"time"
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
	Log.Info("file_reading_started")
	fileData, err := ioutil.ReadFile(configuration.DataPath)
	if err != nil {
		fmt.Print(err)
	}

	var intialWordSet = []string{}
	err = json.Unmarshal(fileData, &intialWordSet)
	if err != nil {
		fmt.Println("error:", err)
	}
	Log.Info("file_reading_done")
	start := time.Now()
	for i := 0; i < len(intialWordSet); i++ {
		bucketData.Insert(intialWordSet[i])
	}
	t := time.Now()
	Log.WithField("time_elapsed", t.Sub(start)).Info("Initialiazing_data_in_ds_done")
}

func GetDs(ctx context.Context) DataStructure {
	if bucketData != nil {
		return bucketData
	}
	return bucketData
}
