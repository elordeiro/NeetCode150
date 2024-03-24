package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}
	buckets := make([][]int, len(nums)+1)
	for num, count := range counts {
		buckets[count] = append(buckets[count], num)
	}
	res := make([]int, 0)
	for i := len(buckets) - 1; i >= 0; i-- {
		if len(buckets[i]) > 0 {
			res = append(res, buckets[i]...)
		}
		if len(res) >= k {
			break
		}
	}
	return res[:k]
}

func print_failed_test(i int, res []int, expected []int) {
	fmt.Printf("Test %d failed:\n", i+1)
	fmt.Printf("    \tActual  : %v\n", res)
	fmt.Printf("    \tExpected: %v\n", expected)
}

func main() {
	tests := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
	}

	passed_all := true
	for i, test := range tests {
		nums, k, expected := test.nums, test.k, test.expected
		res := topKFrequent(nums, k)
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
