package main

import (
	"fmt"
	"math"
)

func maxBananas(piles []int) int {
	max := 0
	for _, bananas := range piles {
		if bananas > max {
			max = bananas
		}
	}
	return max
}

func minEatingSpeed(piles []int, h int) int {
	minK := 1
	maxK := maxBananas(piles)
	k := 0
	for minK <= maxK {
		midK := (minK + maxK) / 2
		hours := 0
		broke := false
		for _, bananas := range piles {
			hours += int(math.Ceil(float64(bananas) / float64(midK)))
			if hours > h {
				broke = true
				break
			}
		}
		if !broke {
			k = midK
			maxK = midK - 1
			continue
		}
		minK = midK + 1
	}
	return k
}

func print_failed_test(i int, res int, expected int) {
	fmt.Printf("Test %d failed:\n", i+1)
	fmt.Printf("    \tActual  : %v\n", res)
	fmt.Printf("    \tExpected: %v\n", expected)
}

func main() {
	tests := []struct {
		piles    []int
		h        int
		expected int
	}{
		{[]int{3, 6, 7, 11}, 8, 4},
		{[]int{30, 11, 23, 4, 20}, 5, 30},
		{[]int{30, 11, 23, 4, 20}, 6, 23},
	}

	passed_all := true
	for i, test := range tests {
		piles, h, expected := test.piles, test.h, test.expected
		res := minEatingSpeed(piles, h)
		if res != expected {
			passed_all = false
			print_failed_test(i, res, expected)
		}
	}
	if passed_all {
		fmt.Println("Passed all tests.")
	}
}
