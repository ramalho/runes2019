// param_map.go
package main

import "fmt"

func removeEmpty(m map[int]string) {
	for k, v := range m {
		if len(v) == 0 {
			delete(m, k)
		}
	}
}

func main() {
	m := map[int]string{0: "", 1: "one", 2: "two", 9: ""}
	removeEmpty(m)
	fmt.Printf("%#v\n", m)
	var m0 map[int]string
	removeEmpty(m0)
	fmt.Printf("%#v\n", m0)
}
