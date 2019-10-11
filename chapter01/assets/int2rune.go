package main

import "fmt"

func main() {
  code := 0x2108
  fmt.Printf("%%T: %T\n", code)
  fmt.Printf("%%v: %v\n", code)	
  fmt.Println("----")

  char := rune(code)
  fmt.Printf("%%T: %T\n", char)	
  fmt.Printf("%%v: %v\n", char)	
  fmt.Printf("%%c: %c\n", char)	
  fmt.Printf("%%q: %q\n", char)	
  fmt.Printf("%%U: %U\n", char)	
  fmt.Printf("%%x: %x\n", char)	
  fmt.Printf("%%d: %d\n", char)	
}

