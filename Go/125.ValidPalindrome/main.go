package main

import (
	"fmt"
	"unicode"
)

func isAlNum(s byte) bool {
	c := rune(s)
	return unicode.IsLetter(c) || unicode.IsNumber(c)
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !isAlNum(s[i]) {
			i++
		}
		for i < j && !isAlNum(s[j]) {
			j--
		}

		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
			return false
		}
		i++
		j--
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
		s        string
		expected bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
	}

	passed_all := true
	for i, test := range tests {
		s, expected := test.s, test.expected
		res := isPalindrome(s)

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
