package main

import (
	"fmt"
	"testing"
)

func Test371SumOfTwoIntegers(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{-1, 1, 0},
		{1, 2, 3},
		{2, 3, 5},
		{3, 4, 7},
	}
	testOnly := 0
	for i, test := range tests {
		if testOnly != 0 && testOnly != i+1 {
			continue
		}
		actual := getSum(test.a, test.b)
		t.Run("Test371SumOfTwoIntegers "+fmt.Sprint(i+1), func(t *testing.T) {
			if actual != test.expected {
				t.Errorf("\n\tActual  : %v\n\tExpected: %v", actual, test.expected)
			}
		})
	}
}
