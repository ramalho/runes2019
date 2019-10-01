package main

import "testing"

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

func Test_scan(t *testing.T) {
	expected := CharName{'\u0020', "SPACE"}
	count := 0
	char := CharName{}
	for char = range scan('\u0000', '\u0021') {
		count++
	}
	if char != expected {
		t.Errorf("expected %q, got %q", expected, char)
	}
	if count != 1 {
		t.Errorf("expected count = 1, got %d", count)
	}
}
