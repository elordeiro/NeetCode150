package main

import (
	"fmt"
	"slices"
)

func wordBreak(s string, wordDict []string) bool {
	n1 := len(s) + 1
	wordSet := make(map[string]struct{})
	for _, word := range wordDict {
		wordSet[word] = struct{}{}
	}

	shortest := len(slices.MinFunc(wordDict, func(a, b string) int { return len(a) - len(b) }))
	longest := len(slices.MaxFunc(wordDict, func(a, b string) int { return len(a) - len(b) }))

	dp := make([]bool, n1)
	dp[0] = true

	for i := shortest; i < n1; i++ {
		for j := i - shortest; j > -1 && j > i-longest-1; j-- {
			if _, ok := wordSet[s[j:i]]; dp[j] && ok {
				dp[i] = true
				break
			}
		}
	}
	return dp[n1-1]
}

func main() {
	tests := []struct {
		s        string
		wordDict []string
		expected bool
	}{
		{"leetcode", []string{"leet", "code"}, true},
		{"applepenapple", []string{"apple", "pen"}, true},
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
		{"bb", []string{"a", "b", "bbb", "bbbb"}, true},
		{"aaaaaaa", []string{"aaaa", "aaa"}, true},
		{"abcd", []string{"a", "abc", "b", "cd"}, true},
	}
	failed := false
	testOnly := 0
	for i, test := range tests {
		if testOnly != 0 && testOnly != i+1 {
			continue
		}
		actual := wordBreak(test.s, test.wordDict)
		if actual != test.expected {
			failed = true
			fmt.Printf("Test %v Failed\n\tActual  : %v\n\tExpected: %v", i+1, actual, test.expected)
		}
	}
	if !failed {
		fmt.Println("All Tests Passed.")
	}
}
