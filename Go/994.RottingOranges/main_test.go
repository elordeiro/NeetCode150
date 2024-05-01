package main

import (
	"fmt"
	"testing"
)

func Test994RottingOranges(t *testing.T) {
	tests := []struct {
		grid     [][]int
		expected int
	}{
		{[][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}, 4},
		{[][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}, -1},
		{[][]int{{0, 2}}, 0},
		{[][]int{{0}}, 0},
		{[][]int{{1}}, -1},
		{[][]int{{2, 2}, {1, 1}, {0, 0}, {2, 0}}, 1},
	}
	for i, test := range tests {
		actual := orangesRotting(test.grid)
		t.Run("Test994RottingOranges"+fmt.Sprint(i+1), func(t *testing.T) {
			if actual != test.expected {
				t.Errorf("\n\tActual  : %v\n\tExpected: %v\n", actual, test.expected)
			}
		})
	}
}
