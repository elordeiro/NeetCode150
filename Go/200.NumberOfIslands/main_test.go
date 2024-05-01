package main

import (
	"fmt"
	"testing"
)

func Test200NumberOfIslands(t *testing.T) {
	tests := []struct {
		grid     [][]byte
		expected int
	}{
		{
			[][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			}, 1,
		},
		{
			[][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			}, 3,
		},
	}
	for i, test := range tests {
		actual := numIslands(test.grid)
		t.Run("Test200NumberOfIslands"+fmt.Sprint(i+1), func(t *testing.T) {
			if actual != test.expected {
				t.Errorf("\n\tActual  : %v\n\tExpected: %v\n", actual, test.expected)
			}
		})
	}
}
