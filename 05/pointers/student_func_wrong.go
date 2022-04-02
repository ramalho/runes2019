// record.go
package main

import "fmt"

type Student struct {
	Name    string
	Credits int
}

func addCredits(s Student, c int) {
	s.Credits += c
}

func main() {
	s := Student{"Anna Raven", 0}
	fmt.Printf("%#v\n", s)
	addCredits(s, 100)
	fmt.Printf("%#v\n", s)
}
