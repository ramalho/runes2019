// change_items_safe.go
package main

import "fmt"

func abbreviate(words []string) []string {
	result := make([]string, len(words)) // HL
	copy(result, words)                  // copy(to, from); to ← from // HL
	for i, w := range result {           // HL
		result[i] = w[:3] // HL
	}
	return result
}

func main() {
	months := []string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	summer := months[6:9]
	q2 := months[4:7]
	q2 = abbreviate(q2)
	fmt.Printf("q2:\t%v\n", q2)
	fmt.Printf("summer:\t%v\n", summer)
	fmt.Printf("months:\t%v\n", months)
}
