// replace.go
package main

import "strings"

type FromTo struct {
	From string
	To   string
}

func multiReplace(text string, changes []FromTo) string { // HL
	for _, c := range changes { // HL
		text = strings.ReplaceAll(text, c.From, c.To) // HL
	} // HL
	return text // HL
} // HL

func main() {
	l33t := []FromTo{{"a", "4"}, {"e", "3"}, {"i", "1"}, {"o", "0"}}
	phrase := "Mad skilled noob powned leet."
	phrase = multiReplace(phrase, l33t) // HL
	println(phrase)
}
