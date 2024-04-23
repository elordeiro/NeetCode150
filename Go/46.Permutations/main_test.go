package main

import (
	"reflect"
	"testing"
)

func Test46Permutations(t *testing.T) {
	tests := []struct {
		nums     []int
		expected [][]int
	}{
		{[]int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
		{[]int{0, 1}, [][]int{{0, 1}, {1, 0}}},
		{[]int{1}, [][]int{{1}}},
	}
	for _, test := range tests {
		actual := permute(test.nums)
		t.Run("Test46Permutations", func(t *testing.T) {
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("Test Failed\n\tActual  : %v\n\tExpected: %v", actual, test.expected)
			}
		})
	}
}
