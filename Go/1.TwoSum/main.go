package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := map[int]int{}

	for i, num := range nums {
		if j, ok := m[target-num]; ok {
			return []int{i, j}
		}
		m[num] = i
	}
	return nil
}

func print_failed_test(i int, res []int, expected []int) {
	fmt.Printf("Test %d failed:\n", i+1)
	fmt.Printf("    \tActual  : %v\n", res)
	fmt.Printf("    \tExpected: %v\n", expected)
}

func main() {
	tests := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	passed_all := true
	for i, test := range tests {
		nums, target, expected := test.nums, test.target, test.expected
		res := twoSum(nums, target)
		if len(res) != len(expected) {
			passed_all = false
			print_failed_test(i, res, expected)
			continue
		}

		if (res[0] != expected[0] || res[1] != expected[1]) &&
			(res[0] != expected[1] || res[1] != expected[0]) {
			passed_all = false
			print_failed_test(i, res, expected)
			break
		}

	}
	if passed_all {
		fmt.Println("Passed all tests.")
	}
}
