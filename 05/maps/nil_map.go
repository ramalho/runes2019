// nil_map.go

package main

import "fmt"

func main() {
	m1 := map[int]string{}
	fmt.Println("// m1 == nil? ", m1 == nil)
	fmt.Printf("m1[3]: %#v\n", m1[3]) // ok
	m1[3] = "three"                   // ok
	fmt.Printf("m1[3]: %#v\n", m1[3]) // ok

	var m0 map[int]string
	fmt.Println("\n// m0 == nil? ", m0 == nil)
	fmt.Printf("m0[3]: %#v\n", m0[3]) // ok
	m0[3] = "three"                   // panic!
}
