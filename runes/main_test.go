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

func Test_scan_just_space(t *testing.T) {
	expected := CharName{'\u0020', "SPACE"}
	count := 0
	cn := CharName{}
	for cn = range scan('\u0020', '\u0021') {
		count++
	}
	if cn != expected {
		t.Errorf("expected %q, got %q", expected, cn)
	}
	if count != 1 {
		t.Errorf("expected count = 1, got %d", count)
	}
}

func Test_scan(t *testing.T) {
	testCases := []struct {
		start rune
		end rune
		expected []CharName
	}{
		{'\x19', '\x21', []CharName{{' ', "SPACE"}}},
		{'0', '3', []CharName{
			{'0', "DIGIT ZERO"},
			{'1', "DIGIT ONE"},
			{'2', "DIGIT TWO"},
		}},
		{'\U000E01EE', '\U000F0000', []CharName{
			{'\U000E01EE', "VARIATION SELECTOR-255"},
			{'\U000E01EF', "VARIATION SELECTOR-256"},
		}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%x:%x", tc.start, tc.end), func(t *testing.T) {
			actual := []CharName{}
			cn := CharName{}
			for cn = range scan(tc.start, tc.end) {
				actual = append(actual, cn)
			}
			assert.Equal(t, tc.expected, actual)
		})
	}
}
