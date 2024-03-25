package main

import "fmt"

func isValid(s string) bool {
	m := map[rune]rune{'(': ')', '[': ']', '{': '}'}
	stack := []rune{}

	for _, char := range s {
		if open, ok := m[char]; ok {
			stack = append(stack, open)
		} else {
			if len(stack) == 0 || char != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func print_failed_test(i int, res bool, expected bool) {
	fmt.Printf("Test %d failed:\n", i+1)
	fmt.Printf("    \tActual  : %v\n", res)
	fmt.Printf("    \tExpected: %v\n", expected)
}

func main() {
	tests := []struct {
		s        string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"[", false},
		{"]", false},
	}

	passed_all := true
	for i, test := range tests {
		s, expected := test.s, test.expected
		res := isValid(s)
		if res != expected {
			passed_all = false
			print_failed_test(i, res, expected)
		}
	}
	if passed_all {
		fmt.Println("Passed all tests.")
	}
}
