package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Example() {
	main()
	// Output:
	// Please provide one or more words to search.
}

func Example_report() {
	report("scruple")
	// Output:
	// U+2108	℈	SCRUPLE
	// 1 character found
}

func Test_CharName_String(t *testing.T) {
	want := "U+0041\tA\tLATIN CAPITAL LETTER A"
	cn := CharName{'A', "LATIN CAPITAL LETTER A"}
	got := fmt.Sprint(cn)
	assert.Equal(t, want, got)
}

func Test_scan_A(t *testing.T) {
	// given
	start := 65
	end := 66
	// expected
	want := []CharName{{'A', "LATIN CAPITAL LETTER A"}}
	got := scan(start, end)
	assert.Equal(t, want, got)
	assert.Equal(t, 1, len(got))
}
