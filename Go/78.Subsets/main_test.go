package main

import (
	"reflect"
	"testing"
)

func Test98Subsets(t *testing.T) {
	tests := []struct {
		nums     []int
		expected [][]int
	}{
		{[]int{1, 2, 3}, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}},
		{[]int{0}, [][]int{{}, {0}}},
	}
	for _, test := range tests {
		actual := subsets(test.nums)
		t.Run("Test98Subsets", func(t *testing.T) {
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("Test Failed\n\tActual  : %v\n\tExpected: %v", actual, test.expected)
			}
		})
	}
}
