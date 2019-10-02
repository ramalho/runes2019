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

func (cn CharName) display() string {
	return fmt.Sprintf("%U\t%c\t%s", cn.Char, cn.Char, cn.Name)
}

func scan(start, end rune) <-chan CharName {
	ch := make(chan CharName)
	go func() {
		for char := start; char < end; char ++ {
			name := runenames.Name(char)
			if name[0] != '<' {
				ch <- CharName{char, name}
			}
		}
		close(ch)
	}()
	return ch
}

func report(words ...string) {
	charName := CharName{'\u2108', "SCRUPLE"}
	fmt.Printf("%s\n", charName.display())
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
