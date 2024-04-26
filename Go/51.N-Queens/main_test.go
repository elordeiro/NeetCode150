package main

import (
	"reflect"
	"testing"
)

func TestSolveNQueens(t *testing.T) {
	tests := []struct {
		n        int
		expected [][]string
	}{
		{4, [][]string{
			{
				".Q..",
				"...Q",
				"Q...",
				"..Q.",
			},
			{
				"..Q.",
				"Q...",
				"...Q",
				".Q..",
			},
		}},
		{1, [][]string{{"Q"}}},
	}
	for i, test := range tests {
		actual := solveNQueens(test.n)
		t.Run("TestSolveNQueens", func(t *testing.T) {
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("Test %v Failed\n\tActual  : %v\n\tExpected: %v", i+1, actual, test.expected)
			}
		})
	}
}
