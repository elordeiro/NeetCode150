package main

import (
	"fmt"
	"slices"
)

func generateParenthesis(n int) []string {
	res := []string{}

	var recur func(opened int, closed int, partial string)
	recur = func(opened int, closed int, partial string) {
		if opened > n {
			return
		}
		if closed == n {
			res = append(res, partial)
			return
		}
		if opened == closed {
			recur(opened+1, closed, partial+"(")
			return
		}
		recur(opened+1, closed, partial+"(")
		recur(opened, closed+1, partial+")")
	}
	recur(0, 0, "")
	return res
}

func print_failed_test(i int, res []string, expected []string) {
	fmt.Printf("Test %d failed:\n", i+1)
	fmt.Printf("    \tActual  : %v\n", res)
	fmt.Printf("    \tExpected: %v\n", expected)
}

func main() {
	tests := []struct {
		n        int
		expected []string
	}{
		{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		{1, []string{"()"}},
	}

	passed_all := true
	for i, test := range tests {
		n, expected := test.n, test.expected
		res := generateParenthesis(n)
		if !slices.Equal(res, expected) {
			passed_all = false
			print_failed_test(i, res, expected)
			break
		}

	}
	if passed_all {
		fmt.Println("Passed all tests.")
	}
}
