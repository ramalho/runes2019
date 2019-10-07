package main

import "fmt"

func main() {
	b := true
	i := 1
	f := 0.1
	c := 3 + 4i
	r := 'A'
	t := "ABC"
	s := []int{10, 20, 30}
	a := [3]float32{}

	fmt.Printf("var b %T = %#v\n", b, b)
	fmt.Printf("var i %T = %#v\n", i, i)
	fmt.Printf("var f %T = %#v\n", f, f)
	fmt.Printf("var c %T = %#v\n", c, c)
	fmt.Printf("var r %T = %#v\n", r, r)
	fmt.Printf("var t %T = %#v\n", t, t)
	fmt.Printf("var s %T = %#v\n", s, s)
	fmt.Printf("var a %T = %#v\n", a, a)
}
