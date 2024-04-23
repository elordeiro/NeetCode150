package main

import "testing"

func Test79WordSearch(t *testing.T) {
	tests := []struct {
		board    [][]byte
		word     string
		expected bool
	}{
		{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED", true},
		{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "SEE", true},
		{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB", false},
		{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED", true},
		{[][]byte{{'a'}}, "a", true},
		{[][]byte{{'A', 'A', 'A', 'A', 'A', 'A'}, {'A', 'A', 'A', 'A', 'A', 'A'}, {'A', 'A', 'A', 'A', 'A', 'A'}, {'A', 'A', 'A', 'A', 'A', 'A'}, {'A', 'A', 'A', 'A', 'A', 'A'}, {'A', 'A', 'A', 'A', 'A', 'A'}}, "AAAAAAAAAAAAAAa", false},
	}
	for i, test := range tests {
		actual := exist(test.board, test.word)
		t.Run("Test79WordSearch", func(t *testing.T) {
			if actual != test.expected {
				t.Errorf("Test %v Failed\n\tActual  : %v\n\tExpected: %v", i+1, actual, test.expected)
			}
		})
	}
}
