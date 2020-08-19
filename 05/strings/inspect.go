package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func inspectString(source, aString string) {

	fmt.Println(source)

	// Make pointer to string structure
	stringPtr := unsafe.Pointer(&aString)
	ptrSize := unsafe.Sizeof(stringPtr)

	// Make pointer to content
	contentPtr := (*byte)(unsafe.Pointer(*(*uintptr)(stringPtr)))

	// Compute address of len field
	lenAddr := uintptr(stringPtr) + ptrSize

	// Make pointer to len field
	lenPtr := (*int)(unsafe.Pointer(lenAddr))

	fmt.Print("\nstruct:")

	fmt.Printf("\t@%p: data %T = %p\n", stringPtr, contentPtr, contentPtr)
	fmt.Printf("\t@%p: len %T = %d\n", lenPtr, *lenPtr, *lenPtr)

	fmt.Print("\ndata:")

	for index := 0; index < len(aString); index++ {
		// Compute address to one byte of content
		byteAddr := uintptr(unsafe.Pointer(contentPtr)) + uintptr(index)
		// Make pointer to that byte of content
		bytePtr := (*byte)(unsafe.Pointer(byteAddr))
		fmt.Printf("\t@%p: %T = %d\n",
			bytePtr, *bytePtr, *bytePtr)
	}
	fmt.Println(strings.Repeat("─", 60))
}

func main() {
	aString := "ABC"
	aStringCopy := aString
	anotherString := "ABC"

	// Same data
	inspectString(`aString := "ABC"`, aString)
	inspectString(`aStringCopy := aString`, aStringCopy)
	inspectString(`anotherString := "ABC"`, anotherString)

	fmt.Println("// updating aString")
	fmt.Println(strings.Repeat("─", 60))

	// Changing data
	inspectString(`aString := "ABC"`, aString)
	aString += "?"
	inspectString(`aString += "?"`, aString)
	aString += "!"
	inspectString(`aString += "!"`, aString)
}
