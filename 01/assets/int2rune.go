package main

import "fmt"

func main() {
	code := 0x20ac
	fmt.Printf("%%T: %T\n", code)
	fmt.Printf("%%v: %v\n", code)
	fmt.Printf("%%x: %x\n", code)
	fmt.Printf("%%d: %d\n", code)
	fmt.Println("----")

	char := '\u20ac'
	fmt.Printf("%%T: %T\n", char)
	fmt.Printf("%%v: %v\n", char)
	fmt.Printf("%%x: %x\n", char)
	fmt.Printf("%%d: %d\n", char)
	fmt.Printf("%%c: %c\n", char)
	fmt.Printf("%%q: %q\n", char)
	fmt.Printf("%%U: %U\n", char)
}
