package main

import (
	"fmt"
)

func countSubstrings(s string) int {
	count := 0
	n := len(s)
	for i := range n {
		for j := range 2 {
			l, r := i, i+j
			for l > -1 && r < n && s[l] == s[r] {
				count++
				l--
				r++
			}
		}
	}
	return count
}

func main() {
	tests := []struct {
		s        string
		expected int
	}{
		{"abc", 3},
		{"aaa", 6},
		{"aba", 4},
		{"dnncbwoneinoplypwgbwktmvkoimcooyiwirgbxlcttgteqthcvyoueyftiwgwwxvxvg", 77},
	}

	failed := false
	testOnly := 0
	for i, test := range tests {
		if testOnly != 0 && testOnly != i+1 {
			continue
		}
		actual := countSubstrings(test.s)
		if actual != test.expected {
			failed = true
			fmt.Printf("Test %v Failed\n\tActual  : %v \n\tExpected: %v\n", i+1, actual, test.expected)
		}
	}
	if !failed {
		fmt.Println("All Tests Passed.")
	}
}
