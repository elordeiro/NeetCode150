package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	haveCount := map[rune]int{}
	needCount := map[rune]int{}
	for _, c := range t {
		needCount[c]++
	}
	minWindow := [2]int{0, 0}
	minWindowLen := math.MaxInt
	l := 0
	have, need := 0, len(needCount)
	for r, char := range s {
		if _, ok := needCount[char]; ok {
			haveCount[char]++

			if haveCount[char] == needCount[char] {
				have++
			}

			for have == need {
				if (r - l + 1) < minWindowLen {
					minWindow = [2]int{l, r + 1}
					minWindowLen = r - l + 1
				}

				if _, ok := needCount[rune(s[l])]; ok {
					haveCount[rune(s[l])]--
					if haveCount[rune(s[l])] < needCount[rune(s[l])] {
						have--
					}
				}
				l++
			}
		}
	}
	return s[minWindow[0]:minWindow[1]]
}

func print_failed_test(i int, res string, expected string) {
	fmt.Printf("Test %d failed:\n", i+1)
	fmt.Printf("    \tActual  : %v\n", res)
	fmt.Printf("    \tExpected: %v\n", expected)
}

func main() {
	tests := []struct {
		s        string
		t        string
		expected string
	}{
		// {"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
		{"a", "aa", ""},
		{"a", "b", ""},
		{"ab", "b", "b"},
		{"aa", "aa", "aa"},
		{"bba", "ab", "ba"},
		{"bbaac", "aba", "baa"},
		// {"aaaaaaaaaaaabbbbbcdd", "abcdd", "abbbbbcdd"},
	}

	passed_all := true
	for i, test := range tests {
		s, t, expected := test.s, test.t, test.expected
		res := minWindow(s, t)
		if res != expected {
			passed_all = false
			print_failed_test(i, res, expected)

		}
	}
	if passed_all {
		fmt.Println("Passed all tests.")
	}
}
