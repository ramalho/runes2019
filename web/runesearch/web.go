package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/runenames"
)

const hostAddr = "localhost:8000"

// CharName holds a single Unicode character and its name
type CharName struct {
	Char string `json:"char"`
	Name string `json:"name"`
}

var index = buildIndex(scan(0, unicode.MaxRune))

func main() {
	fmt.Printf("Serving on: http://%s", hostAddr)
	http.HandleFunc("/words", WordSearch)
	http.HandleFunc("/", Form)
	log.Fatal(http.ListenAndServe(hostAddr, nil))
}

// Form serves the HTML form
func Form(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "form.html")
}

func makeResults(chars []rune) []CharName {
	result := []CharName{}
	for _, char := range chars {
		result = append(result, CharName{
			Char: string(char),
			Name: runenames.Name(char),
		})
	}
	return result
}

// WordSearch returns characters that have query words in their names
func WordSearch(w http.ResponseWriter, req *http.Request) {
	chars := []rune{}
	query := strings.TrimSpace(req.URL.Query().Get("q"))
	if len(query) > 0 {
		chars = Search(index, query)
	}
	js, err := json.Marshal(makeResults(chars))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(js)
}
