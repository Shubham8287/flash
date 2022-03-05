package hashmap

import (
	"container/heap"
	"fmt"
)

type Item struct {
	query string
	freq string
}

type PriorityQueue []*Item

map[string]PriorityQueue