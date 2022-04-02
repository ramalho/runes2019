func Test_scan(t *testing.T) {
	testCases := []struct {
		label string
		start rune
		end   rune
		want  []CharName
	}{
		{"A", 'A', 'B', []CharName{{'A', "LATIN CAPITAL LETTER A"}}},
		// inputs/outputs for other test cases here...
	}
	for _, tc := range testCases {
		t.Run(tc.label, func(t *testing.T) {
			got := scan(tc.start, tc.end)
			assert.Equal(t, tc.want, got)
		})
	}
}