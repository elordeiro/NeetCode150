package main

import (
	"fmt"
	"testing"
)

func Test136SingleNumber(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
		{[]int{1}, 1},
	}
	for i, test := range tests {
		actual := singleNumber(test.nums)
		t.Run("Test136SingleNumber "+fmt.Sprint(i+1), func(t *testing.T) {
			if actual != test.expected {
				t.Errorf("\n\tActual : %v\b\n\tExpected: %v\n", actual, test.expected)
			}
		})
	}
}
