package main

import "fmt"

func function(nums []int) []int {
	num := 1
	res := []int{}

	for i := range len(nums) {
		res = append(res, num)
		num *= nums[i]
	}

	num = 1
	for i := len(nums) - 1; i > -1; i-- {
		res[i] *= num
		num *= nums[i]
	}

	return res
}

func print_failed_test(i int, res []int, expected []int) {
	fmt.Printf("Test %d failed:\n", i+1)
	fmt.Printf("    \tActual  : %v\n", res)
	fmt.Printf("    \tExpected: %v\n", expected)
}

func main() {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
	}

	passed_all := true
	for i, test := range tests {
		nums, expected := test.nums, test.expected
		res := function(nums)
		if len(res) != len(expected) {
			passed_all = false
			print_failed_test(i, res, expected)
			continue
		}
		for j, num := range res {
			if num != expected[j] {
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
