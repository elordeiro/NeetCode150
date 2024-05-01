package main

func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var dfs func(int, int) int
	dfs = func(row, col int) int {
		if row < 0 || col < 0 || row == m || col == n || grid[row][col] < 1 {
			return 0
		}
		grid[row][col] = -1
		return (1 +
			dfs(row+1, col) +
			dfs(row, col+1) +
			dfs(row-1, col) +
			dfs(row, col-1))
	}
	largest := 0
	for row := range m {
		for col := range n {
			if grid[row][col] == 1 {
				largest = max(largest, dfs(row, col))
			}
		}
	}
	return largest
}
