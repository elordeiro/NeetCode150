package main

func hammingWeight(n int) int {
	total := 0
	for n > 0 {
		total += 1 & n
		n >>= 1
	}
	return total
}
