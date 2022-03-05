package trie

func (t *trie) generateAllWordsWithPrefix(currentNode *trieNode, currentString string) []string {
	matchingStrings := []string{}
	usingDFS(currentNode, &matchingStrings, currentString)
	return matchingStrings
}

func usingDFS(currentNode *trieNode, matchingStrings *[]string, currentString string) {
	if currentNode.isWordEnd {
		*matchingStrings = append(*matchingStrings, (currentString))
	}
	for i := 0; i < 26; i++ {
		if currentNode.childrens[i] != nil {
			usingDFS(currentNode.childrens[i], matchingStrings, currentString+(string(i+int('a'))))
		}
	}
}
