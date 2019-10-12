package main

import "fmt"

func main() {
	s := "Erick"
	b := s[0]
	fmt.Printf("%q\t%d\n", b, b)
	s = "Érico"
	b = s[0]
	fmt.Printf("%q\t%d\n", b, b)
	runes := []rune(s)
	c := runes[0]
	fmt.Printf("%q\t%d\n", c, c)
}
