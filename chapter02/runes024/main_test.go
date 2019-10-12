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
	start, end := 'A', 'B'
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

func Test_scan_ABC(t *testing.T) {
	// given
	start, end := 'A', 'D'
	// when
	got := scan(start, end)
	// expected
	want := []CharName{
		{'A', "LATIN CAPITAL LETTER A"},
		{'B', "LATIN CAPITAL LETTER B"},
		{'C', "LATIN CAPITAL LETTER C"},
	}
	assert.Equal(t, want, got)
}

func Test_scan_unassigned(t *testing.T) {
	// given
	start, end := '\u0377', '\u037B'
	// when
	got := scan(start, end)
	// expected
	want := []CharName{
		{'\u0377', "GREEK SMALL LETTER PAMPHYLIAN DIGAMMA"},
		{'\u037A', "GREEK YPOGEGRAMMENI"},
	}
	assert.Equal(t, want, got)
}

func Test_scan_unnamed(t *testing.T) {
	// given
	start, end := '\x1E', '\x22'
	// when
	got := scan(start, end)
	// expected
	want := []CharName{
		{' ', "SPACE"},
		{'!', "EXCLAMATION MARK"},
	}
	assert.Equal(t, want, got)
}
