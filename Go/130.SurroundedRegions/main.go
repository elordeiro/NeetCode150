package main

import (
	"fmt"
	"reflect"
)

func solve(board [][]byte) {
	m, n := len(board)-1, len(board[0])-1

	var dfs func(int, int)
	dfs = func(row, col int) {
		if board[row][col] == 'O' {
			board[row][col] = 'C'
			if row > 0 {
				dfs(row-1, col)
			}
			if row < m {
				dfs(row+1, col)
			}
			if col > 0 {
				dfs(row, col-1)
			}
			if col < n {
				dfs(row, col+1)
			}
		}
	}

	for row := range m + 1 {
		if board[row][0] == 'O' {
			dfs(row, 0)
		}
		if board[row][n] == 'O' {
			dfs(row, n)
		}
	}
	for col := range n + 1 {
		if board[0][col] == 'O' {
			dfs(0, col)
		}
		if board[m][col] == 'O' {
			dfs(m, col)
		}
	}

	for row := range m + 1 {
		for col := range n + 1 {
			if board[row][col] == 'C' {
				board[row][col] = 'O'
			} else if board[row][col] == 'O' {
				board[row][col] = 'X'
			}
		}
	}
}

func main() {
	tests := []struct {
		board    [][]byte
		expected [][]byte
	}{
		{[][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}, [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'O', 'X', 'X'}}},
		{[][]byte{{'X'}}, [][]byte{{'X'}}},
		{[][]byte{{'O', 'X', 'X', 'O', 'X'}, {'X', 'O', 'O', 'X', 'O'}, {'X', 'O', 'X', 'O', 'X'}, {'O', 'X', 'O', 'O', 'O'}, {'X', 'X', 'O', 'X', 'O'}}, [][]byte{{'O', 'X', 'X', 'O', 'X'}, {'X', 'X', 'X', 'X', 'O'}, {'X', 'X', 'X', 'O', 'X'}, {'O', 'X', 'O', 'O', 'O'}, {'X', 'X', 'O', 'X', 'O'}}},
	}
	testOnly := 0
	for i, test := range tests {
		if testOnly != 0 && testOnly != i+1 {
			continue
		}
		solve(test.board)
		if !reflect.DeepEqual(test.board, test.expected) {
			fmt.Printf("Test %v Failed\n\tActual  : %v\n\tExpected: %v", test.board, test.expected)
		}
	}
}
