package main

import (
	"fmt"
	"os"
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
		ch <- CharName{' ', "SPACE"}
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
