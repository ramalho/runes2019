// replace_in_place.go
package main

import "strings"

type FromTo struct {
	From string
	To   string
}

func multiReplaceInPlace(textPtr *string, changes []FromTo) { // HL
	for _, c := range changes { // HL
		*textPtr = strings.ReplaceAll(*textPtr, c.From, c.To) // HL
	} // HL
} // HL

func main() {
	l33t := []FromTo{{"a", "4"}, {"e", "3"}, {"i", "1"}, {"o", "0"}}
	phrase := "Mad skilled noob powned leet."
	multiReplaceInPlace(&phrase, l33t) // HL
	println(phrase)
}
