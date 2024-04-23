package main

import "slices"

func permute(nums []int) [][]int {
	n := len(nums)
	res := [][]int{}
	var backtrack func(int, []int, []int)
	backtrack = func(i int, partial, curr []int) {
		if len(partial) == n {
			res = append(res, partial)
			return
		}
		if i >= len(curr) {
			return
		}
		backtrack(0, slices.Concat(partial, []int{curr[i]}), slices.Concat(curr[:i], curr[i+1:]))
		backtrack(i+1, partial, curr)
	}
	backtrack(0, []int{}, nums)
	return res
}
