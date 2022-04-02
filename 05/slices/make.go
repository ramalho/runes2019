package main

import "fmt"

func main() {
	var s0 []int
	s1 := []int{}
	s2 := make([]int, 0)
	s3 := make([]int, 3)
	s4 := make([]int, 4, 8)
	for _, s := range [][]int{s0, s1, s2, s3, s4} {
		fmt.Printf("%T len=%d cap%d\t%#v\n", s, len(s), cap(s), s)
	}
}
