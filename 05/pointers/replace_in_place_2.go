// replace_in_place_2.go
package main

import "strings"

type FromTo struct {
	From string
	To   string
}

func multiReplaceInPlace(textPtr *string, changes []FromTo) {
	for _, c := range changes {
		*textPtr = strings.ReplaceAll(*textPtr, c.From, c.To)
	}
}

func main() {
	l33t := []FromTo{{"a", "4"}, {"e", "3"}, {"i", "1"}, {"o", "0"}}
	phrase := "Mad skilled noob powned leet."
	slogan := phrase                   // HL
	multiReplaceInPlace(&slogan, l33t) // HL
	println(phrase)                    // HL
	println(slogan)                    // HL
}
