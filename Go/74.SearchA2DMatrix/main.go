package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	u, d := 0, len(matrix)-1
	l, r := 0, len(matrix[0])-1
	found := false
	var m int
	for u <= d {
		m = (u + d) / 2
		if target < matrix[m][0] {
			d = m - 1
		} else if matrix[m][r] < target {
			u = m + 1
		} else {
			found = true
			break
		}
	}
	if !found {
		return false
	}

	row := m
	for l <= r {
		m := (l + r) / 2
		if target < matrix[row][m] {
			r = m - 1
		} else if matrix[row][m] < target {
			l = m + 1
		} else {
			return true
		}
	}
	return false
}

func print_failed_test(i int, res bool, expected bool) {
	fmt.Printf("Test %d failed:\n", i+1)
	fmt.Printf("    \tActual  : %v\n", res)
	fmt.Printf("    \tExpected: %v\n", expected)
}

func main() {
	tests := []struct {
		matrix   [][]int
		target   int
		expected bool
	}{
		{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3, true},
		{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13, false},
	}

	passed_all := true
	for i, test := range tests {
		matrix, target, expected := test.matrix, test.target, test.expected
		res := searchMatrix(matrix, target)
		if res != expected {
			passed_all = false
			print_failed_test(i, res, expected)
		}
	}
	if passed_all {
		fmt.Println("Passed all tests.")
	}
}
