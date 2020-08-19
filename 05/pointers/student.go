// student.go
package main

import "fmt"

type Student struct {
	Name    string
	Credits int
}

func main() {
	s := Student{"Anna Raven", 270}
	fmt.Printf("s:  %T\t%#v\n", s, s)
	var p *Student
	// fmt.Println(*p) panic: nil pointer dereference
	fmt.Printf("p:  %T\t%#v\n", p, p)
	p = &s // the address of s
	fmt.Printf("p:  %T\t%#v\n", p, p)
	a := *p // the Student referenced by p
	fmt.Printf("a:  %T\t%#v\n", a, a)
	fmt.Printf("*p: %T\t%#v\n", *p, *p)
}
