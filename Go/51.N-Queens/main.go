package main

var queens []int

func createBoard() []string {
	n := len(queens)
	board := []string{}
	for r := range n {
		row := ""
		for c := range n {
			if c == queens[r] {
				row += "Q"
			} else {
				row += "."
			}
		}
		board = append(board, row)
	}
	return board
}

func solveNQueens(n int) [][]string {
	res := [][]string{}
	queens = make([]int, n)
	cols := make([]bool, n)
	hills := make([]bool, n*4)
	dales := make([]bool, n*2)
	var backtrack func(int)
	backtrack = func(row int) {
		if row == n {
			res = append(res, createBoard())
			return
		}
		for col := range n {
			if !cols[col] && !hills[row-col+2*n] && !dales[row+col] {
				queens[row] = col
				cols[col] = true
				hills[row-col+2*n] = true
				dales[row+col] = true
				backtrack(row + 1)
				queens[row] = -1
				cols[col] = false
				hills[row-col+2*n] = false
				dales[row+col] = false
			}
		}
	}
	backtrack(0)
	return res
}
