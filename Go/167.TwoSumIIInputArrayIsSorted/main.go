package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	var total int
	for {
		total = numbers[i] + numbers[j]
		if total < target {
			i++
		} else if total > target {
			j--
		} else {
			return []int{i + 1, j + 1}
		}
	}
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
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
		{[]int{-1, 0}, -1, []int{1, 2}},
		{[]int{3, 24, 50, 79, 88, 150, 345}, 200, []int{3, 6}},
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
