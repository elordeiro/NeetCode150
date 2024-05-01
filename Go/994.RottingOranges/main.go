package main

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	queue := [][2]int{}
	toRot := make(map[[2]int]struct{})

	for row := range m {
		for col := range n {
			if grid[row][col] == 1 {
				toRot[[2]int{row, col}] = struct{}{}
			} else if grid[row][col] == 2 {
				toRot[[2]int{row, col}] = struct{}{}
				queue = append(queue, [2]int{row, col})
			}
		}
	}

	if len(toRot) == 0 {
		return 0
	}

	dirs := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	time := -1
	idx := 0
	for idx < len(queue) {
		time++
		for range len(queue) - idx {
			row, col := queue[idx][0], queue[idx][1]
			delete(toRot, queue[idx])
			idx++
			for _, dir := range dirs {
				r, c := row+dir[0], col+dir[1]
				if r > -1 && r < m && c > -1 && c < n && grid[r][c] == 1 {
					grid[r][c] = 2
					queue = append(queue, [2]int{r, c})
				}
			}

		}
	}
	if len(toRot) > 0 {
		return -1
	}
	return time
}
