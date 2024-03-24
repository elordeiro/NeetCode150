package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[rune]int)
	for _, c := range s {
		m[c]++
	}
	for _, c := range t {
		m[c]--
	}
	for _, v := range m {
		if v != 0 {
			return false
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
		s        string
		t        string
		expected bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"a", "ab", false},
	}

	passed_all := true
	for i, test := range tests {
		s, t, expected := test.s, test.t, test.expected
		res := isAnagram(s, t)
		if res != expected {
			passed_all = false
			print_failed_test(i, res, expected)

		}
	}
	if passed_all {
		fmt.Println("Passed all tests.")
	}
}
