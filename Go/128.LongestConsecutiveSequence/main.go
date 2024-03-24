package main

import "fmt"

func longestConsecutive(nums []int) int {
	m := make(map[int]struct{})
	for _, num := range nums {
		m[num] = struct{}{}
	}

	longest := 0
	for _, curr := range nums {

		delete(m, curr)
		length := 1

		next := curr + 1
		for _, ok := m[next]; ok; _, ok = m[next] {
			delete(m, next)
			next++
			length++
		}

		prev := curr - 1
		for _, ok := m[prev]; ok; _, ok = m[prev] {
			delete(m, prev)
			prev--
			length++
		}

		longest = max(longest, length)
	}
	return longest
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
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
		{[]int{1, 0, -1, 2}, 4},
		{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 1},
	}

	passed_all := true
	for i, test := range tests {
		nums, expected := test.nums, test.expected
		res := longestConsecutive(nums)

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
