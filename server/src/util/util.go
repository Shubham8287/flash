package util

type Response struct {
    Prefix string `json:"prefix"`
    Matches []string `json:"matches"`
}

// TODO: sort by frequency
func FilterMatchingStrings(wordList []string) []string {
	if (len(wordList) < 20) {
		return wordList;
	}
	return wordList[0:20];
}