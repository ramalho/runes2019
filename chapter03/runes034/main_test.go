package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

func Example() {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"", "EIGHTHS", "fraction"}
	main()
	// Output:
	// U+215C	⅜	VULGAR FRACTION THREE EIGHTHS
	// U+215D	⅝	VULGAR FRACTION FIVE EIGHTHS
	// U+215E	⅞	VULGAR FRACTION SEVEN EIGHTHS
	// 3 character(s) found
}

func Example_one_arg() {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"", "cruzeiro"}
	main()
	// Output:
	// U+20A2	₢	CRUZEIRO SIGN
	// 1 character(s) found
}

func Example_no_args() {
	main()
	// Output:
	// Please provide one or more words to search.
}

func Example_report() {
	report("scruple")
	// Output:
	// U+2108	℈	SCRUPLE
	// 1 character(s) found
}

func Example_report_2_results() {
	report("copyright")
	// Output:
	// U+00A9	©	COPYRIGHT SIGN
	// U+2117	℗	SOUND RECORDING COPYRIGHT
	// 2 character(s) found
}

func Test_CharName_String(t *testing.T) {
	want := "U+0041\tA\tLATIN CAPITAL LETTER A"
	cn := CharName{'A', "LATIN CAPITAL LETTER A"}
	got := fmt.Sprint(cn)
	assert.Equal(t, want, got)
}

func Test_scan(t *testing.T) {
	testCases := []struct {
		label string
		start rune
		end   rune
		want  []CharName
	}{
		{"A", 'A', 'B', []CharName{{'A', "LATIN CAPITAL LETTER A"}}},
		{"ABC", 'A', 'D', []CharName{
			{'A', "LATIN CAPITAL LETTER A"},
			{'B', "LATIN CAPITAL LETTER B"},
			{'C', "LATIN CAPITAL LETTER C"},
		}},
		{"unassigned", '\u0377', '\u037B', []CharName{
			{'\u0377', "GREEK SMALL LETTER PAMPHYLIAN DIGAMMA"},
			{'\u037A', "GREEK YPOGEGRAMMENI"},
		}},
		{"unnamed", '\x1E', '\x22', []CharName{
			{' ', "SPACE"},
			{'!', "EXCLAMATION MARK"},
		}},
	}
	for _, tc := range testCases {
		t.Run(tc.label, func(t *testing.T) {
			got := scan(tc.start, tc.end)
			assert.Equal(t, tc.want, got)
		})
	}
}

func Test_filter(t *testing.T) {
	testCases := []struct {
		start rune
		end   rune
		query []string
		want  []CharName
	}{
		{' ', unicode.MaxRune, []string{"MADEUPWORD"}, []CharName{}},
		{'\u2108', unicode.MaxRune, []string{"SCRUPLE"}, []CharName{
			{'\u2108', "SCRUPLE"},
		}},
		{'6', '9', []string{"SEVEN"}, []CharName{
			{'7', "DIGIT SEVEN"},
		}},
		{'A', 'C', []string{"A"}, []CharName{
			{'A', "LATIN CAPITAL LETTER A"},
		}},
		{',', '/', []string{"MINUS"}, []CharName{
			{'-', "HYPHEN-MINUS"},
		}},
		{'?', 'C', []string{"LeTtEr", "capital"}, []CharName{
			{'A', "LATIN CAPITAL LETTER A"},
			{'B', "LATIN CAPITAL LETTER B"},
		}},
	}
	for _, tc := range testCases {
		t.Run(strings.Join(tc.query, "+"), func(t *testing.T) {
			sample := scan(tc.start, tc.end)
			got := filter(sample, tc.query)
			assert.Equal(t, tc.want, got)
		})
	}
}
