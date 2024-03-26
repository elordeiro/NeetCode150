package main

import "fmt"

func search(nums []int, target int) int {
	var binSearch func(left, right int) int
	binSearch = func(left, right int) int {
		if nums[left] <= nums[right] {
			for left <= right {
				mid := (left + right) / 2
				if target < nums[mid] {
					right = mid - 1
				} else if nums[mid] < target {
					left = mid + 1
				} else {
					return mid
				}
			}
			return -1
		}

		mid := (left + right) / 2
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target <= nums[mid] {
				return binSearch(left, mid)
			} else {
				return binSearch(mid+1, right)
			}
		}

		if nums[mid] <= target && target <= nums[right] {
			return binSearch(mid, right)
		} else {
			return binSearch(left, mid-1)
		}
	}
	return binSearch(0, len(nums)-1)
}

func print_failed_test(i int, res int, expected int) {
	fmt.Printf("Test %d failed:\n", i+1)
	fmt.Printf("    \tActual  : %v\n", res)
	fmt.Printf("    \tExpected: %v\n", expected)
}

func main() {
	tests := []struct {
		nums     []int
		target   int
		expected int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{1}, 0, -1},
		{[]int{3, 1}, 0, -1},
		{[]int{5, 1, 3}, 1, 1},
	}

	passed_all := true
	for i, test := range tests {
		nums, target, expected := test.nums, test.target, test.expected
		res := search(nums, target)

		if res != expected {
			passed_all = false
			print_failed_test(i, res, expected)
		}
	}
	if passed_all {
		fmt.Println("Passed all tests.")
	}
}
