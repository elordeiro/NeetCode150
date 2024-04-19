package main

import "slices"

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	n := len(candidates) - 1
	var recur func(int, int, []int)
	recur = func(i, partialSum int, partialArr []int) {
		if partialSum == target {
			res = append(res, slices.Clone(partialArr))
			return
		}
		if partialSum > target || i > n {
			return
		}
		recur(i, partialSum+candidates[i], append(partialArr, candidates[i]))
		recur(i+1, partialSum, partialArr)
	}
	recur(0, 0, []int{})
	return res
}
