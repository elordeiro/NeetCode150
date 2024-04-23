package main

import "slices"

func combinationSum2(candidates []int, target int) [][]int {
	slices.Sort(candidates)
	res := [][]int{}
	var backtrack func(int, int, []int)
	backtrack = func(idx int, total int, partial []int) {
		if total == target {
			res = append(res, partial)
			return
		}

		for i, n := idx, len(candidates); i < n; i++ {
			x := candidates[i]
			if x+total > target || x > target {
				break
			}
			if i > idx && x == candidates[i-1] {
				continue
			}
			backtrack(i+1, total+x, slices.Concat(partial, []int{x}))
		}
	}
	backtrack(0, 0, []int{})
	return res
}
