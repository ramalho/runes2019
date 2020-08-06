package main

import (
	"fmt"
	"testing"
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
	if got != want {
		t.Errorf("CharName_String\n\tgot:  %q\n\twant: %q", got, want)
	}
}

func Test_scan_A(t *testing.T) {
	// given
	start, end := 65, 66
	// when
	got := scan(start, end)
	// expected
	want := []CharName{{'A', "LATIN CAPITAL LETTER A"}}
	if len(got) != 1 {
		t.Errorf("CharName_String: len(got) = %d; want 1", len(got))
	}
	if got[0] != want[0] {
		t.Errorf("CharName_String\n\tgot:  %q\n\twant: %q", got, want)
	}
}
