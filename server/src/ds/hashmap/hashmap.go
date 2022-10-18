package hashmap

import (
	"flash/util"
)

type bucketType map[string][]Item

var bucket bucketType = nil

// sample bucket
// [a:[{2 aaa} {2 aa} {1 a} {1 aah} {1 aahed} {1 aahs} {1 aal} {1 aalii} {1 aaliis}]
// aa:[[{2 aaa} {2 aa} {1 aah} {1 aahed} {1 aahs} {1 aal} {1 aalii} {1 aaliis}]
// aaa:[{2 aaa}] aah:[{1 aah} {1 aahed} {1 aahs}] aahe:[{1 aahed}]
// aahed:[{1 aahed}] aahs:[{1 aahs}] aal:[{1 aal} {1 aalii} {1 aaliis}]
// aali:[{1 aalii} {1 aaliis}] aalii:[{1 aalii} {1 aaliis}]
// aaliis:[{1 aaliis}]
// w:[{4 waskodimagama} {3 wask}]
// wa:[{4 waskodimagama} {3 wask}]
// was:[{4 waskodimagama} {3 wask}]
// wask:[{4 waskodimagama} {3 wask}]
// wasko:[{4 waskodimagama}]
// waskod:[{4 waskodimagama}]
// waskodi:[{4 waskodimagama}]
// waskodim:[{4 waskodimagama}]
// waskodima:[{4 waskodimagama}]
// waskodimag:[{4 waskodimagama}]
// waskodimaga:[{4 waskodimagama}]
// waskodimagam:[{4 waskodimagama}]
// waskodimagama:[{4 waskodimagama}]
// ]
type Item struct {
	freq int
	word string
}

func Get() *bucketType {
	if bucket == nil {
		bucket = make(bucketType)
	}
	return &bucket
}

func (b *bucketType) Find(prefix string) []string {
	return filterMatchingWordItem(bucket[prefix])
}

func (b *bucketType) Insert(word string) {
	// sanitize the input string
	word = util.SanitizeString(word)

	var prefixWord string
	for i := 0; i < len(word); i++ {
		prefixWord += string(word[i])
		_, ok := bucket[prefixWord]
		if ok {
			idx, ok := findWord(bucket[prefixWord], word)
			if ok {
				newWordItem := bucket[prefixWord][idx]
				bucket[prefixWord] = removeWordItem(bucket[prefixWord], idx)
				newWordItem.freq++
				bucket[prefixWord] = insertWordItem(bucket[prefixWord], newWordItem)
			} else {
				freq := 1
				if len(bucket[prefixWord]) != 0 {
					freq = bucket[prefixWord][len(bucket[prefixWord])-1].freq
				}
				newWordItem := Item{
					word: word,
					freq: freq,
				}
				bucket[prefixWord] = insertWordItem(bucket[prefixWord], newWordItem)
			}
		} else {
			bucket[prefixWord] = append(bucket[prefixWord], Item{
				word: word,
				freq: 1,
			})
		}
	}
}

func filterMatchingWordItem(wordList []Item) []string {
	var matchingStrings []string
	for i := 0; i < len(wordList) && i < 20; i++ {
		matchingStrings = append(matchingStrings, wordList[i].word)
	}
	return matchingStrings
}

func findWord(wordList []Item, word string) (int, bool) {
	for i := 0; i < len(wordList); i++ {
		if wordList[i].word == word {
			return i, true
		}
	}
	return -1, false
}

func removeWordItem(wordList []Item, idx int) []Item {
	return append(wordList[:idx], wordList[idx+1:]...)
}

func insertWordItem(wordList []Item, newWord Item) []Item {
	if len(wordList) == 100 {
		wordList = wordList[:99]
	}
	pos := len(wordList)
	for i := len(wordList) - 1; i >= 0; i-- {
		if wordList[i].freq >= newWord.freq {
			break
		}
		pos--
	}
	wordList = append(wordList, newWord)
	copy(wordList[pos+1:], wordList[pos:])
	wordList[pos] = newWord
	return wordList

}
