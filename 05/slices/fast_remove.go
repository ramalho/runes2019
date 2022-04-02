// fast_remove.go
package main

import "fmt"

func removeEmptyInPlace(strings []string) []string {
	result := strings[:0] // zero-length slice of original // HL
	for _, s := range strings {
		if s != "" {
			result = append(result, s)
		}
	}
	return result // change in place, but must return new slice struct // HL
}

func main() {
	words := []string{1: "alpha", 3: "charlie", 5: "echo"}
	fmt.Printf("%#v\n", words)
	words = removeEmptyInPlace(words) // 🚨 must reassign to update `words` slice! // HL
	fmt.Printf("%#v\n", words)
}
