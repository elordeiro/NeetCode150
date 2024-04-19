package main

import (
	"reflect"
	"testing"
)

func Test39CombinationSum(t *testing.T) {
	tests := []struct {
		candidates []int
		target     int
		expected   [][]int
	}{
		{[]int{2, 3, 6, 7}, 7, [][]int{{2, 2, 3}, {7}}},
		{[]int{2, 3, 5}, 8, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
		{[]int{2}, 1, [][]int{}},
	}
	testOnly := 2
	for i, test := range tests {
		if testOnly != 0 && i+1 != testOnly {
			continue
		}
		actual := combinationSum(test.candidates, test.target)
		t.Run("Test39CombinationSum", func(t *testing.T) {
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("Test Failed\n\tActual  : %v\n\tExpected: %v", actual, test.expected)
			}
		})
	}
}
