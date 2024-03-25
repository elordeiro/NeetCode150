package main

import "fmt"

func function(nums []int) []int {
	return nums
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
		{[]int{1, 2, 3, 1}, []int{1, 2, 3, 1}},
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}},
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
