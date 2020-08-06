package main

import (
	"fmt"
	"testing"
)

func TestFactorial(t *testing.T) {
	want := int(120)
	got := Factorial(5)
	if got != want {
		t.Errorf("Factorial(5) = %d; want %d", got, want)
	}
}

func TestFactorial_table(t *testing.T) {
	testCases := []struct {
		n    int
		want int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{5, 120},
		{10, 3628800},
		{20, 2432902008176640000},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d", tc.n), func(t *testing.T) {
			got := Factorial(tc.n)
			if got != tc.want {
				t.Errorf("Factorial(%d) = %d; want %d", tc.n, got, tc.want)
			}
		})
	}
}
