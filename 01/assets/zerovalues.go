package main

import "fmt"

func main() {
	var b bool
	var i int
	var f float64
	var c complex128
	var t string
	var s []int // slices are NOT flat
	var a [3]float32

	fmt.Printf("var b %T = %#v\n", b, b)
	fmt.Printf("var i %T = %#v\n", i, i)
	fmt.Printf("var f %T = %#v\n", f, f)
	fmt.Printf("var c %T = %#v\n", c, c)
	fmt.Printf("var t %T = %#v\n", t, t)
	fmt.Printf("var s %T = %#v  // slice has a pointer!\n", s, s)
	fmt.Printf("var a %T = %#v\n", a, a)
}
