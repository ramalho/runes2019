package main

import "fmt"

// Factorial returns n!
func Factorial(n int) int {
	if n < 2 {
		return 1
	}
	result := n
	for n > 2 {
		n--
		result *= n
	}
	return result
}

func main() {
	n := 20                // int is the default type of integer literals
	result := Factorial(n) // result will match the return type of Factorial
	fmt.Println("Largest factorial that fits an int:")
	fmt.Printf("\t%d! = %d\n", n, result)
}
