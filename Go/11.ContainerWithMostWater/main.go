package main

import "fmt"

func maxHeight(height []int) int {
	maxHeight := 0
	for _, h := range height {
		if h > maxHeight {
			maxHeight = h
		}
	}
	return maxHeight
}

func maxArea(height []int) int {
	maxArea := 0
	maxHeight := maxHeight(height)
	i, j := 0, len(height)-1
	var area int
	for i < j {
		area = min(height[i], height[j]) * (j - i)

		if area > maxArea {
			maxArea = area
		}

		if area >= maxHeight*(j-i) {
			break
		}

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return maxArea
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
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 16},
		{[]int{1, 2, 4, 3}, 4},
		{[]int{2, 3, 4, 5, 18, 17, 6}, 17},
	}

	passed_all := true
	for i, test := range tests {
		nums, expected := test.nums, test.expected
		res := maxArea(nums)

		if res != expected {
			passed_all = false
			print_failed_test(i, res, expected)
		}
	}
	if passed_all {
		fmt.Println("Passed all tests.")
	}
}
