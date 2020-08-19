// months.go
package main

import "fmt"

func main() {
	months := []string{
		1: "January", 2: "February", 3: "March", 4: "April",
		5: "May", 6: "June", 7: "July", 8: "August",
		9: "September", 10: "October", 11: "November", 12: "December",
	}
	summer := months[6:9]
	q2 := months[4:7]
	fmt.Printf("months: len=%d cap=%d\t%#v\n", len(months), cap(months), months)
	fmt.Printf("summer: len=%d cap=%d\t%#v\n", len(summer), cap(summer), summer)
	fmt.Printf("q2:     len=%d cap=%d\t%#v\n", len(q2), cap(q2), q2)
}
