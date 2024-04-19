package main

var res [][]int
var globalNums []int

func backtrack(i int, partial []int) {
	if i == -1 {
		res = append(res, partial)
		return
	}
	backtrack(i-1, partial)
	backtrack(i-1, append([]int{globalNums[i]}, partial...))
}

func subsets(nums []int) [][]int {
	res = [][]int{}
	globalNums = nums
	backtrack(len(nums)-1, []int{})
	return res
}
