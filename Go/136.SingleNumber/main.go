package main

func singleNumber(nums []int) int {
	accum := 0
	for _, num := range nums {
		accum ^= num
	}
	return accum
}
