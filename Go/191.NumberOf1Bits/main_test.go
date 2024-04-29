package main

import (
	"fmt"
	"testing"
)

func Test191NumberOf1Bits(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{11, 3},
		{128, 1},
		{2147483645, 30},
	}
	for i, test := range tests {
		actual := hammingWeight(test.n)
		t.Run("Test191NumberOf1Bits "+fmt.Sprint(i+1), func(t *testing.T) {
			if actual != test.expected {
				t.Errorf("\n\tActual  : %v\n\tExpected: %v\n", actual, test.expected)
			}
		})
	}

}
