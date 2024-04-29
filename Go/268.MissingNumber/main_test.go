package main

import (
	"fmt"
	"testing"
)

func Test268MissingNumber(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{3, 0, 1}, 2},
		{[]int{0, 1}, 2},
		{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
		{[]int{19, 58, 31, 51, 54, 47, 68, 25, 85, 9, 83, 70, 24, 75, 30, 78, 62, 38, 41, 21, 56, 60, 94, 1, 45, 15, 72, 52, 28, 93, 14, 96, 35, 17, 95, 89, 74, 46, 13, 82, 57, 76, 55, 20, 36, 63, 44, 61, 6, 92, 65, 50, 91, 42, 98, 34, 8, 33, 40, 12, 7, 48, 11, 80, 10, 71, 97, 39, 73, 26, 99, 43, 90, 5, 3, 2, 23, 29, 0, 79, 53, 64, 4, 27, 37, 84, 69, 81, 22, 86, 67, 32, 66, 18, 16, 77, 59, 49, 87}, 88},
	}
	testOnly := 4
	for i, test := range tests {
		if testOnly != 0 && testOnly != i+1 {
			continue
		}
		actual := missingNumber(test.nums)
		t.Run("Test268MissingNumber "+fmt.Sprint(i+1), func(t *testing.T) {
			if actual != test.expected {
				t.Errorf("\n\tActual  : %v\n\tExpected: %v", actual, test.expected)
			}
		})
	}
}
