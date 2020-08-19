//get_element.go
package main

import "fmt"

func main() {
	m := map[int]string{2: "two", 9: ""}
	elem := m[2]
	fmt.Printf("m[2]: %#v\n", elem)
	_, exists := m[2]
	fmt.Printf("key 2 exists? %#v\n", exists)

	fmt.Println("\nThe “comma ok” idiom:")
	elem, ok := m[3]
	fmt.Printf("m[3]: %#v, %#v\n", elem, ok)
	elem, ok = m[9]
	fmt.Printf("m[9]: %#v, %#v\n", elem, ok)

	fmt.Println("\nDeleting an element:")
	delete(m, 2)
	elem, ok = m[2]
	fmt.Printf("m[2]: %#v, %#v\n", elem, ok)
}
