package trie

import "flash/util"

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

func InitTrie() *trie {
    return &trie{
        root: &trieNode{},
    }
}

func (t *trie) Insert(word string) {
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
    matchingStrings := []string{}
    wordLength := len(word)
    currentNode := t.root
    for i := 0; i < wordLength; i++ {
        index := word[i] - 'a'
        if currentNode.childrens[index] == nil {
            t.Insert(word);
            return matchingStrings;
        }
        currentNode = currentNode.childrens[index]
    }

    t.findAllStringsWithDFS(currentNode, &matchingStrings, word);
    response := util.FilterMatchingStrings(matchingStrings);

	return response
}