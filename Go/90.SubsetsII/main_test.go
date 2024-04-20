package main

import (
	"reflect"
	"testing"
)

func Test90SubsetsII(t *testing.T) {
	tests := []struct {
		nums     []int
		expected [][]int
	}{
		{[]int{1, 2, 2}, [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}}},
		{[]int{0}, [][]int{{}, {0}}},
		{[]int{4, 4, 4, 1, 4}, [][]int{{}, {1}, {1, 4}, {1, 4, 4}, {1, 4, 4, 4}, {1, 4, 4, 4, 4}, {4}, {4, 4}, {4, 4, 4}, {4, 4, 4, 4}}},
	}
	for _, test := range tests {
		actual := subsetsWithDup(test.nums)
		t.Run("Test90SubsetsII", func(t *testing.T) {
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("Test Failed\n\tActual  : %v\n\tExpected: %v", actual, test.expected)
			}
		})
	}
}
