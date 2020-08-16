package main

import (
	"fmt"
	"testing"
	"unicode"

	"github.com/standupdev/runeset"
	"github.com/stretchr/testify/assert"
)

func Test_scan(t *testing.T) {
	testCases := []struct {
		label string
		start rune
		end   rune
		want  []RuneName
	}{
		{"A", 'A', 'B', []RuneName{{'A', "LATIN CAPITAL LETTER A"}}},
		{"ABC", 'A', 'D', []RuneName{
			{'A', "LATIN CAPITAL LETTER A"},
			{'B', "LATIN CAPITAL LETTER B"},
			{'C', "LATIN CAPITAL LETTER C"},
		}},
		{"unassigned", '\u0377', '\u037B', []RuneName{
			{'\u0377', "GREEK SMALL LETTER PAMPHYLIAN DIGAMMA"},
			{'\u037A', "GREEK YPOGEGRAMMENI"},
		}},
		{"unnamed", '\x1E', '\x22', []RuneName{
			{' ', "SPACE"},
			{'!', "EXCLAMATION MARK"},
		}},
	}
	for _, tc := range testCases {
		t.Run(tc.label, func(t *testing.T) {
			got := []RuneName{}
			for cn := range scan(tc.start, tc.end) {
				got = append(got, cn)
			}
			assert.Equal(t, tc.want, got)
		})
	}
}

func Test_tokenize(t *testing.T) {
	var testCases = []struct {
		name string
		want []string
	}{
		{"EXCLAMATION MARK",
			[]string{"EXCLAMATION", "MARK"}},
		{"HYPHEN-MINUS",
			[]string{"HYPHEN", "MINUS"}},
		{"",
			[]string{}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tokenize(tc.name)
			assert.Equal(t, tc.want, got)
		})
	}
}

func Test_buildIndex_twoLines(t *testing.T) {
	// 003D;EQUALS SIGN;Sm;0;ON;;;;;N;;;;;
	// 003E;GREATER-THAN SIGN;Sm;0;ON;;;;;Y;;;;;
	want := Index{
		"EQUALS":  runeset.Make('='),
		"GREATER": runeset.Make('>'),
		"THAN":    runeset.Make('>'),
		"SIGN":    runeset.Make('=', '>'),
	}
	got := buildIndex(scan(0x3D, 0x3F))
	assert.Equal(t, want, got)
}

func Test_buildIndex_threeLines(t *testing.T) {
	// 0041;LATIN CAPITAL LETTER A;Lu;0;L;;;;;N;;;;0061;
	// 0042;LATIN CAPITAL LETTER B;Lu;0;L;;;;;N;;;;0062;
	// 0043;LATIN CAPITAL LETTER C;Lu;0;L;;;;;N;;;;0063;
	want := Index{
		"A":       runeset.Make('A'),
		"B":       runeset.Make('B'),
		"C":       runeset.Make('C'),
		"LATIN":   runeset.MakeFromString("ABC"),
		"CAPITAL": runeset.MakeFromString("ABC"),
		"LETTER":  runeset.MakeFromString("ABC"),
	}
	got := buildIndex(scan(0x41, 0x44))
	assert.Equal(t, want, got)
}

func Test_buildIndex_all(t *testing.T) {
	index := buildIndex(scan(0, unicode.MaxRune))
	const wantAtLeastWords = 10_000
	assert.Greater(t, len(index), wantAtLeastWords,
		fmt.Sprintf("index should have more than %d keys", wantAtLeastWords))
	const registeredSign = '\u00AE' // ®
	wantSet := runeset.Make(registeredSign)
	gotSet := index["REGISTERED"]
	assert.Equal(t, wantSet, gotSet, "REGISTERED should map only to U+00AE")
}

var fullIndex = buildIndex(scan(0, unicode.MaxRune))

func TestSearch(t *testing.T) {
	var testCases = []struct {
		query string
		want  runeset.Set
	}{
		{"Registered", runeset.Make('®')},
		{"ORDINAL", runeset.Make('ª', 'º')},
		{"fraction eighths", runeset.Make('⅜', '⅝', '⅞')},
		{"fraction eighths bang", runeset.Set{}},
		{"fraction eighths five", runeset.Make('⅝')},
		{"NoSuchRune", runeset.Set{}},
		{"", runeset.Set{}},
	}
	for _, tc := range testCases {
		t.Run(tc.query, func(t *testing.T) {
			got := Search(fullIndex, tc.query)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestSearch_hyphenatedQuery(t *testing.T) {
	query := "HYPHEN-MINUS"
	want := '-'
	got := Search(fullIndex, query)
	if !got.Contains(want) {
		t.Errorf("query: %q\t%q absent, got: %v",
			query, want, got)
	}
}
