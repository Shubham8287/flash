package trie

import (
	"flash/util"
)

const (
	//ALBHABET_SIZE total characters in english alphabet
	ALBHABET_SIZE = 26
)

type trieNode struct {
	childrens [ALBHABET_SIZE]*trieNode
	isWordEnd bool
}

type trie struct {
	root *trieNode
}

var trieInstance *trie = nil

func Get() *trie {
	if trieInstance == nil {
		trieInstance = &trie{
			root: &trieNode{},
		}
	}
	return trieInstance
}

func (t *trie) Insert(word string) {
	// sanitize the input string
	word = util.SanitizeString(word)
	
	wordLength := len(word)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if currentNode.childrens[index] == nil {
			currentNode.childrens[index] = &trieNode{}
		}
		currentNode = currentNode.childrens[index]
	}
	currentNode.isWordEnd = true
}

func (t *trie) Find(word string) []string {
	wordLength := len(word)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if currentNode.childrens[index] == nil {
			//t.Insert(word)
			return []string{}
		}
		currentNode = currentNode.childrens[index]
	}

	matchingStrings := t.generateAllWordsWithPrefix(currentNode, word)
	response := util.FilterMatchingStrings(matchingStrings)

	return response
}
