package main

func missingNumber(nums []int) int {
	n := len(nums)
	expect := n * (n + 1) / 2
	total := 0
	for _, num := range nums {
		total += num
	}
	return expect - total
}
