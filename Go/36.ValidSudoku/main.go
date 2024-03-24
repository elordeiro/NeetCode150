package main

import "fmt"

type Set map[byte]struct{}

func (s *Set) add(item byte) *Set {
	if s == nil {
		s := make(Set)
		s[item] = struct{}{}
		return &s
	}
	(*s)[item] = struct{}{}
	return s
}

func (s *Set) remove(item byte) {
	delete(*s, item)
}

func (s *Set) contains(item byte) bool {
	if s == nil {
		return false
	}
	_, found := (*s)[item]
	return found
}

func isValidSudoku(board [][]byte) bool {
	rows := make([]Set, 9)
	cols := make([]Set, 9)
	sqrs := make(map[[2]int]*Set)

	for i := range 9 {
		rows[i] = make(Set)
		cols[i] = make(Set)
	}

	for i := range 9 {
		for j := range 9 {
			cell := board[i][j]
			if cell == '.' {
				continue
			}
			sqr := [2]int{i / 3, j / 3}
			if rows[i].contains(cell) ||
				cols[j].contains(cell) ||
				sqrs[sqr].contains(cell) {
				return false
			}
			rows[i].add(cell)
			cols[j].add(cell)
			sqrs[sqr] = sqrs[sqr].add(cell)
		}
	}
	return true
}

func print_failed_test(i int, res bool, expected bool) {
	fmt.Printf("Test %d failed:\n", i+1)
	fmt.Printf("    \tActual  : %v\n", res)
	fmt.Printf("    \tExpected: %v\n", expected)
}

func main() {
	tests := []struct {
		board    [][]byte
		expected bool
	}{
		{
			[][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			}, true,
		},
		{
			[][]byte{
				{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			}, false,
		},
		{
			[][]byte{
				{'.', '.', '.', '.', '5', '.', '.', '1', '.'},
				{'.', '4', '.', '3', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '3', '.', '.', '1'},
				{'8', '.', '.', '.', '.', '.', '.', '2', '.'},
				{'.', '.', '2', '.', '7', '.', '.', '.', '.'},
				{'.', '1', '5', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '2', '.', '.', '.'},
				{'.', '2', '.', '9', '.', '.', '.', '.', '.'},
				{'.', '.', '4', '.', '.', '.', '.', '.', '.'},
			}, false,
		},
		{
			[][]byte{
				{'.', '.', '4', '.', '.', '.', '6', '3', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'5', '.', '.', '.', '.', '.', '.', '9', '.'},
				{'.', '.', '.', '5', '6', '.', '.', '.', '.'},
				{'4', '.', '3', '.', '.', '.', '.', '.', '1'},
				{'.', '.', '.', '7', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '5', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
			}, false,
		},
	}

	passed_all := true
	test_only := 4
	for i, test := range tests {
		if test_only != 0 && test_only != i+1 {
			continue
		}
		board, expected := test.board, test.expected
		res := isValidSudoku(board)
		if res != expected {
			passed_all = false
			print_failed_test(i, res, expected)
			break
		}

	}

	if passed_all {
		fmt.Println("Passed all tests.")
	}
}
