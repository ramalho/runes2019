// safe_remove.go
package main

import "fmt"

func removeEmpty(strings []string) []string {
	var result []string // new slice, not derived from strings // HL
	for _, s := range strings {
		if s != "" {
			result = append(result, s)
		}
	}
	return result // return new slice // HL
}

func main() {
	words := []string{1: "alpha", 3: "charlie", 5: "echo"}
	fmt.Printf("%#v\n", words)
	words = removeEmpty(words) // reassign to get new slice // HL
	fmt.Printf("%#v\n", words)
}
