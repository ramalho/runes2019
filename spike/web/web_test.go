package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/standupdev/runeset"
	"github.com/stretchr/testify/assert"
)

func Test_makeResults(t *testing.T) {
	var testCases = []struct {
		given runeset.Set
		want  []CharName
	}{
		{runeset.Set{}, []CharName{}},
		{runeset.Make('2', '0', '3', '1'), []CharName{
			{"0", "DIGIT ZERO"},
			{"1", "DIGIT ONE"},
			{"2", "DIGIT TWO"},
			{"3", "DIGIT THREE"},
		}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprint(tc.given), func(t *testing.T) {
			got := makeResults(tc.given)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestWordSearch(t *testing.T) {
	var testCases = []struct {
		query string
		want  string
	}{
		{"scruple", `[{"char": "℈", "name": "SCRUPLE"}]`},
		{"Indicator Ordinal",
			`[{"char":"ª", "name":"FEMININE ORDINAL INDICATOR"},
			  {"char":"º", "name":"MASCULINE ORDINAL INDICATOR"}]`},
		{"", `[]`},
		{"NoSuchWordExists", `[]`},
	}
	for _, tc := range testCases {
		t.Run(tc.query, func(t *testing.T) {
			url := fmt.Sprintf("any-url/word?q=%s", tc.query)
			request, err := http.NewRequest("GET", url, nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}
			recorder := httptest.NewRecorder()
			WordSearch(recorder, request)
			response := recorder.Result()
			assert.Equal(t, http.StatusOK, response.StatusCode)
			wantContentType := "application/json; charset=utf-8"
			assert.Equal(t, wantContentType, response.Header.Get("Content-type"))
			defer response.Body.Close()
			content, err := ioutil.ReadAll(response.Body)
			if err != nil {
				t.Fatalf("could not read response body: %v", err)
			}
			assert.JSONEq(t, tc.want, string(content))
		})
	}
}
