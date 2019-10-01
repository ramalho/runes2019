package main

import (
	"fmt"
	"os"
	"golang.org/x/text/unicode/runenames"
)

type CharName struct {
	Char rune
	Name string
}


func scan(start, end rune) <-chan CharName {
	ch := make(chan CharName)
	go func() {
		ch <- CharName{' ', "SPACE"}
		close(ch)
	}()
	return ch
}

func report(words ...string) {
	char := '\u2108'
	name := runenames.Name(char)
	fmt.Printf("%U\t%c\t%s\n", char, char, name)
	count := 1
	fmt.Printf("%d character found", count)
}

func main() {
	if len(os.Args) > 1 {
		report(os.Args[1:]...)
	} else {	
		fmt.Println("Please provide one or more words to search.")
	}
}
