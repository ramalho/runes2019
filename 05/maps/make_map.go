// make_map.go

package main

import "fmt"

func main() {
	var m0 map[int]string
	m1 := map[int]string{}
	m2 := make(map[int]string)
	m3 := make(map[int]string, 10_000)
	for _, s := range []map[int]string{m0, m1, m2, m3} {
		isNil := s == nil
		fmt.Printf("len: %d  nil: %-5v  %#v\n", len(s), isNil, s)
	}
}
