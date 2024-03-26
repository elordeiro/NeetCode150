package main

import "fmt"

func findMin(nums []int) int {
	l, r := 0, len(nums)-1

	if nums[l] <= nums[r] {
		return nums[l]
	}

	mid := (l + r) / 2
	return min(findMin(nums[:mid+1]), findMin(nums[mid+1:]))
}

func print_failed_test(i int, res int, expected int) {
	fmt.Printf("Test %d failed:\n", i+1)
	fmt.Printf("    \tActual  : %v\n", res)
	fmt.Printf("    \tExpected: %v\n", expected)
}

func main() {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{3, 4, 5, 1, 2}, 1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
		{[]int{11, 13, 15, 17}, 11},
		{[]int{2, 1}, 1},
		{[]int{3, 1, 2}, 1},
	}

	passed_all := true
	for i, test := range tests {
		nums, expected := test.nums, test.expected
		res := findMin(nums)
		if res != expected {
			passed_all = false
			print_failed_test(i, res, expected)
		}
	}
	if passed_all {
		fmt.Println("Passed all tests.")
	}
}
