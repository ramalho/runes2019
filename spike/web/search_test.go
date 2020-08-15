package main

import (
	"testing"
	"unicode"

	"github.com/standupdev/runeset"
	"github.com/stretchr/testify/assert"
)

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
