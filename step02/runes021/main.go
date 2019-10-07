/* step01/runes/main.go */
package main

import (
	"fmt"

	"golang.org/x/text/unicode/runenames"
)

type CharName struct {
	Char rune
	Name string
}

func (c CharName) String() string {
	return fmt.Sprintf("%U\t%c\t%s", c.Char, c.Char, c.Name)
}

func report(word string) {
	char := '\u2108'
	name := runenames.Name(char)
	fmt.Printf("%U\t%c\t%s\n", char, char, name)
	count := 1
	fmt.Printf("%d character found", count)
}

func main() {
	fmt.Println("Please provide one or more words to search.")
}
