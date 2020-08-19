// record.go
package main

import "fmt"

type Student struct {
	Name    string
	Credits int
}

func addCredits(s *Student, c int) { // HL
	s.Credits += c
}

func main() {
	s := Student{"Anna Raven", 0}
	fmt.Printf("%#v\n", s)
	addCredits(&s, 100) // HL
	fmt.Printf("%#v\n", s)
}
