package main

import "fmt"

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, num := range nums {
		if m[num] {
			return true
		}
		m[num] = true
	}
	return false
}

func print_failed_test(i int, res bool, expected bool) {
	fmt.Printf("Test %d failed:\n", i+1)
	fmt.Printf("    \tActual  : %v\n", res)
	fmt.Printf("    \tExpected: %v\n", expected)
}

func main() {
	tests := []struct {
		nums     []int
		expected bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	passed_all := true
	for i, test := range tests {
		nums, expected := test.nums, test.expected
		res := containsDuplicate(nums)
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
