package trie

func (t *trie) findAllStringsWithDFS(currentNode *trieNode, matchingStrings *[]string, currentString string) {
    if currentNode.isWordEnd {
        *matchingStrings = append(*matchingStrings, (currentString));
    }
    for i:=0; i<26; i++ {
        if currentNode.childrens[i] != nil {
            t.findAllStringsWithDFS(currentNode.childrens[i], matchingStrings, currentString+(string(i+int('a'))));
        }    
    }
}