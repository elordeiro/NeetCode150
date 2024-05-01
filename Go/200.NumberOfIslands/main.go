package main

func numIslands(grid [][]byte) int {
	m, n := len(grid)-1, len(grid[0])-1
	var dfs func(int, int)
	dfs = func(row, col int) {
		if grid[row][col] == '1' {
			grid[row][col] = '#'
			if row > 0 {
				dfs(row-1, col)
			}
			if col > 0 {
				dfs(row, col-1)
			}
			if row < m {
				dfs(row+1, col)
			}
			if col < n {
				dfs(row, col+1)
			}
		}
	}
	islands := 0
	for row := range m + 1 {
		for col := range n + 1 {
			if grid[row][col] == '1' {
				islands++
				dfs(row, col)
			}
		}
	}
	return islands
}
