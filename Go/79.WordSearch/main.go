package main

func allCharsInBoard(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	for _, c := range word {
		char := byte(c)
		found := false
		for i := range m {
			for j := range n {
				if char == board[i][j] {
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func exist(board [][]byte, word string) bool {
	if !allCharsInBoard(board, word) {
		return false
	}
	m, n, len_word := len(board), len(board[0]), len(word)

	var backtrack func(int, int, int) bool
	backtrack = func(row int, col int, i int) bool {
		if i == len_word {
			return true
		}
		if row < 0 || row >= m || col < 0 || col >= n || word[i] != board[row][col] || board[row][col] == '#' {
			return false
		}
		old := board[row][col]
		board[row][col] = '#'
		res := (backtrack(row+1, col, i+1) ||
			backtrack(row-1, col, i+1) ||
			backtrack(row, col+1, i+1) ||
			backtrack(row, col-1, i+1))
		board[row][col] = old
		return res
	}
	for i := range m {
		for j := range n {
			if backtrack(i, j, 0) {
				return true
			}
		}
	}
	return false
}
