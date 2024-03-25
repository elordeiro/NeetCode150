package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	l, r, longest := 0, 0, 0
	substring := make(map[byte]struct{})
	for r < n {
		c := s[r]
		if _, ok := substring[c]; !ok {
			substring[c] = struct{}{}
		} else {
			longest = max(longest, len(substring))
			for s[l] != c {
				delete(substring, s[l])
				l++
			}
			l++
		}
		r++
	}
	longest = max(longest, len(substring))
	return longest
}

func print_failed_test(i int, res int, expected int) {
	fmt.Printf("Test %d failed:\n", i+1)
	fmt.Printf("    \tActual  : %v\n", res)
	fmt.Printf("    \tExpected: %v\n", expected)
}

func main() {
	tests := []struct {
		nums     string
		expected int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{" ", 1},
		{"dvdf", 3},
	}

	passed_all := true
	for i, test := range tests {
		s, expected := test.nums, test.expected
		res := lengthOfLongestSubstring(s)
		if res != expected {
			passed_all = false
			print_failed_test(i, res, expected)
		}
	}
	if passed_all {
		fmt.Println("Passed all tests.")
	}
}
