package wordcounter

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
)

// dictionaryItem is an item in the word store
type dictionaryItem struct {
	Word     string
	Count    int
	Previous *dictionaryItem
	Next     *dictionaryItem
}

// dictionary is the word store
var dictionary *dictionaryItem

// init initialises the word store
func init() {
	reset()
}

// reset resets the word store to an empty state
func reset() {
	dictionary = &dictionaryItem{}
}

// response writes a successful response as JSON to the client
func response(w http.ResponseWriter, res interface{}) {
	if err, ok := res.(error); ok {
		ErrorResponse(w, err)
		return
	}

	resBody, err := json.Marshal(res)
	if err != nil {
		ErrorResponse(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resBody)
}

// ErrorResponse writes out an error to the client as plaintext
func ErrorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(err.Error()))
}

func (d *dictionaryItem) isEmpty() bool {
	return d == nil || d.Word == ""
}

// counts returns a map of all words and their counts in the word store
func (d *dictionaryItem) counts() map[string]int {
	if d.isEmpty() {
		return nil
	}

	result := d.Previous.counts()
	if result == nil {
		result = map[string]int{d.Word: d.Count}
	} else {
		result[d.Word] = d.Count
	}

	remaining := d.Next.counts()
	for k, v := range remaining {
		result[k] = v
	}

	return result
}

// GetCounts returns words in the word store with their counts as JSON to the client
func GetCounts(w http.ResponseWriter) {
	response(w, dictionary.counts())
}

// ProcessInput takes the POST payload and processes it into the word store
func ProcessInput(w http.ResponseWriter, r *http.Request) {
	count := 0
	input := bufio.NewScanner(r.Body)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		dictionary.add(rtrim(input.Text()))
		count++
	}

	response(w, fmt.Sprint("Words processed: ", count))
}

// rtrim removes a non-alpha character from the end of the word, e.g. punctuation (assumes max one non-alpha character)
func rtrim(word string) string {
	last := len(word) - 1
	char := word[last]
	if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
		return word
	}
	return word[:last]
}

// add implements a binary tree for the word store
func (d *dictionaryItem) add(word string) {
	if d.Word == "" {
		d.Word = word
		d.Count = 1
		return
	}

	if d.Word == word {
		d.Count++
		return
	}

	if d.Word < word {
		if d.Next == nil {
			d.Next = &dictionaryItem{Word: word, Count: 1}
		} else {
			d.Next.add(word)
		}
	} else {
		if d.Previous == nil {
			d.Previous = &dictionaryItem{Word: word, Count: 1}
		} else {
			d.Previous.add(word)
		}
	}
}
