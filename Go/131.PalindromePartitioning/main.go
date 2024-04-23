package main

import "slices"

var word string

func isPalindrome(i, j int) bool {
	for i < j {
		if word[i] != word[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func partition(s string) [][]string {
	n := len(s)
	word = s
	res := [][]string{}
	partial := []string{}
	var backtrack func(int)
	backtrack = func(i int) {
		if i >= n {
			res = append(res, slices.Clone(partial))
			return
		}
		for j := i; j < n; j++ {
			if isPalindrome(i, j) {
				partial = append(partial, s[i:j+1])
				backtrack(j + 1)
				partial = partial[:len(partial)-1]

			}
		}
	}
	backtrack(0)
	return res
}
