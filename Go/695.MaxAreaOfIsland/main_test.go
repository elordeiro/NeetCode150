package main

import (
	"fmt"
	"testing"
)

func Test695MaxAreaOfIsland(t *testing.T) {
	tests := []struct {
		grid     [][]int
		expected int
	}{
		{[][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}, 6},
		{[][]int{{0, 0, 0, 0, 0, 0, 0, 0}}, 0},
	}
	for i, test := range tests {
		actual := maxAreaOfIsland(test.grid)
		t.Run("Test695MaxAreaOfIsland"+fmt.Sprint(i+1), func(t *testing.T) {
			if actual != test.expected {
				t.Errorf("\n\tActual  : %v\n\tExpected: %v", actual, test.expected)
			}
		})
	}
}
