package wordcounter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

// NewWordStore initialises the word store
func NewWordStore() {
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
func GetCounts(w http.ResponseWriter, r *http.Request) {
	response(w, dictionary.counts())
}

// ProcessInput takes the POST payload and processes it into the word store
func ProcessInput(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	defer r.Body.Close()

	// step through characters in the body building up words and adding them to the word store
	count := 0
	var b bytes.Buffer
	for _, char := range body {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			b.WriteString(string(char))
		} else {
			if word := b.String(); word != "" {
				dictionary.add(word)
				count++
				b.Reset()
			}
		}
	}

	response(w, fmt.Sprint("Words processed: ", count))
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
