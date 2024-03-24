package main

import (
	"fmt"
	"slices"
)

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	n := len(nums)
	i, j, k := 0, 0, 0
	slices.Sort(nums)
	for i < n {
		target := -nums[i]
		j, k = i+1, n-1
		for j < k {
			total := nums[j] + nums[k]
			if total < target {
				j++
			} else if total > target {
				k--
			} else {
				res = append(res, []int{nums[i], nums[j], nums[k]})

				prev := nums[j]
				for j < n && prev == nums[j] {
					j++
				}

				prev = nums[k]
				for k > j && prev == nums[k] {
					k--
				}
			}
		}
		prev := nums[i]
		for i < n && prev == nums[i] {
			i++
		}
	}
	return res
}

func print_failed_test(i int, res [][]int, expected [][]int) {
	fmt.Printf("Test %d failed:\n", i+1)
	fmt.Printf("    \tActual  : %v\n", res)
	fmt.Printf("    \tExpected: %v\n", expected)
}

func main() {
	tests := []struct {
		nums     []int
		expected [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{0, 1, 1}, [][]int{{}}},
		{[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
	}

	passed_all := true
	for i, test := range tests {
		nums, expected := test.nums, test.expected
		res := threeSum(nums)
		for j, arr := range res {
			if !slices.Equal(arr, expected[j]) {
				passed_all = false
				print_failed_test(i, res, expected)
				break
			}
		}
	}
	if passed_all {
		fmt.Println("Passed all tests.")
	}
}
