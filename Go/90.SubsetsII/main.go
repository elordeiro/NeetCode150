package main

import (
	"fmt"
	"slices"
)

func subsetsWithDup(nums []int) [][]int {
	slices.Sort(nums)
	res := [][]int{}
	var backtrack func([]int, []int)
	backtrack = func(partial, avail []int) {
		if len(avail) == 0 {
			res = append(res, partial)
			return
		}

		backtrack(slices.Concat(partial, []int{avail[0]}), avail[1:])

		i := 1
		for i < len(avail) && avail[0] == avail[i] {
			i++
		}
		backtrack(partial, avail[i:])
	}
	backtrack([]int{}, nums)
	return res
}

func main() {
	arr := []int{1}
	fmt.Println(arr)
	fmt.Println(arr[:])
	fmt.Println(arr[:1])
	fmt.Println(arr[1:])
}
