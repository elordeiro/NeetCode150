package main

import "fmt"

func trap(height []int) int {
	var n int
	if n = len(height); n < 3 {
		return 0
	}
	water, totalWater := 0, 0
	l, r := 0, 1
	for r < n {
		if height[r] < height[l] {
			water += height[l] - height[r]
		} else {
			totalWater += water
			water = 0
			l = r
		}
		r++
	}
	n = l
	r--
	l = r - 1
	for l > n {
		if height[l] < height[r] {
			totalWater += height[r] - height[l]
		} else {
			r = l
		}
		l--
	}
	return totalWater
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
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{[]int{4, 2, 0, 3, 2, 5}, 9},
	}

	passed_all := true
	for i, test := range tests {
		nums, expected := test.nums, test.expected
		res := trap(nums)

		if res != expected {
			passed_all = false
			print_failed_test(i, res, expected)
			break
		}
	}

	if passed_all {
		fmt.Println("Passed all tests.")
	}
}
