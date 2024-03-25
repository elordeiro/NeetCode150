package main

import "fmt"

func characterReplacement(s string, k int) int {
	l, r, mostFreq := 0, 0, 0
	m := make(map[byte]int)

	for r = range len(s) {
		m[s[r]]++
		mostFreq = max(mostFreq, m[s[r]])

		if (r - l + 1 - mostFreq) > k {
			m[s[l]]--
			l++
		}
	}
	return r - l + 1
}

func print_failed_test(i int, res int, expected int) {
	fmt.Printf("Test %d failed:\n", i+1)
	fmt.Printf("    \tActual  : %v\n", res)
	fmt.Printf("    \tExpected: %v\n", expected)
}

func main() {
	tests := []struct {
		s        string
		k        int
		expected int
	}{
		{"ABAB", 2, 4},
		{"AABABBA", 1, 4},
		{"AABALMNOPQRS", 1, 4},
	}

	passed_all := true
	for i, test := range tests {
		s, k, expected := test.s, test.k, test.expected
		res := characterReplacement(s, k)
		if res != expected {
			passed_all = false
			print_failed_test(i, res, expected)

		}
	}
	if passed_all {
		fmt.Println("Passed all tests.")
	}
}
