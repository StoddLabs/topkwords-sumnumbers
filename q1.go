package cos418_hw1_1

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

// Find the top K most common words in a text document.
//
//	path: location of the document
//	numWords: number of words to return (i.e. k)
//	charThreshold: character threshold for whether a token qualifies as a word,
//		e.g. charThreshold = 5 means "apple" is a word but "pear" is not.
//
// Matching is case insensitive, e.g. "Orange" and "orange" is considered the same word.
// A word comprises alphanumeric characters only. All punctuations and other characters
// are removed, e.g. "don't" becomes "dont".
// You should use `checkError` to handle potential errors.
func topWords(path string, numWords int, charThreshold int) []WordCount {
	// TODO: implement me
	// HINT: You may find the `strings.Fields` and `strings.ToLower` functions helpful
	// HINT: To keep only alphanumeric characters, use the regex "[^0-9a-zA-Z]+"

	f, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("oh no")
	}
	words := strings.Fields(string(f))

	//wm will keep the position of where the word is in ws(the list)
	//ws keeps each unique word, and using wm, it can quickly find the word, and increment the count
	wm := make(map[string]int)
	ws := make([]WordCount, 0)

	cursor := 0

	for _, v := range words {
		//we don't care about case/non-alphanumeric character
		w := strings.ToLower(v)
		w = regexp.MustCompile("[^0-9a-zA-Z]+").ReplaceAllString(w, "")
		//caller and decide if the length of the word matter or not
		if charThreshold > len(w) {
			continue
		}
		if _, ok := wm[w]; !ok {
			ws = append(ws, WordCount{w, 1})
			wm[w] = cursor
			cursor += 1
		} else {
			i := wm[w]
			ws[i].Count += 1
		}
	}
	sortWordCounts(ws)

	tr := ws[0:numWords]

	return tr
}

// A struct that represents how many times a word is observed in a document
type WordCount struct {
	Word  string
	Count int
}

func (wc WordCount) String() string {
	return fmt.Sprintf("%v: %v", wc.Word, wc.Count)
}

// Helper function to sort a list of word counts in place.
// This sorts by the count in decreasing order, breaking ties using the word.
// DO NOT MODIFY THIS FUNCTION!
func sortWordCounts(wordCounts []WordCount) {
	sort.Slice(wordCounts, func(i, j int) bool {
		wc1 := wordCounts[i]
		wc2 := wordCounts[j]
		if wc1.Count == wc2.Count {
			return wc1.Word < wc2.Word
		}
		return wc1.Count > wc2.Count
	})
}
