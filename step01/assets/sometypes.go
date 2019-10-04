package main

import "fmt"

func main() {
	i := 1
	f := 0.1
	c := 'A'
	t := "ABC"
	s := []int{10, 20, 30}

	fmt.Printf("var i %T = %#v\n", i, i)
	fmt.Printf("var f %T = %#v\n", f, f)
	fmt.Printf("var c %T = %#v\n", c, c)
	fmt.Printf("var t %T = %#v\n", t, t)
	fmt.Printf("var s %T = %#v\n", s, s)
}

