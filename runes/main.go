package main

import (
	"fmt"
	"os"
	"golang.org/x/text/unicode/runenames"
)

func display(words ...string) {
	char := '\u2108'
	name := runenames.Name(char)
	fmt.Printf("%U\t%c\t%s\n", char, char, name)
	count := 1
	fmt.Printf("%d character found", count)
}

func main() {
	if len(os.Args) > 1 {
		display(os.Args[1:]...)
	} else {	
		fmt.Println("Please provide one or more words to search.")
	}
}
