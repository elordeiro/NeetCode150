package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	half := (m + n) / 2
	if m > n {
		nums1, nums2 = nums2, nums1
		m, n = n, m
	}
	l, r := 0, m-1
	for {
		midA := int(math.Floor(float64(l+r) / float64(2)))
		midB := half - midA - 2
		var leftA, leftB, rightA, rightB int
		if midA > -1 {
			leftA = nums1[midA]
		} else {
			leftA = math.MinInt
		}
		if midB > -1 {
			leftB = nums2[midB]
		} else {
			leftB = math.MinInt
		}

		if midA+1 < m {
			rightA = nums1[midA+1]
		} else {
			rightA = math.MaxInt
		}
		if midB+1 < n {
			rightB = nums2[midB+1]
		} else {
			rightB = math.MaxInt
		}

		if leftA <= rightB && leftB <= rightA {
			if (m+n)%2 == 0 {
				return float64(max(leftA, leftB)+min(rightA, rightB)) / 2.0
			}
			return float64(min(rightA, rightB))
		}
		if leftA > rightB {
			r = midA - 1
		} else {
			l = midA + 1
		}
	}
}

func print_failed_test(i int, res float64, expected float64) {
	fmt.Printf("Test %d failed:\n", i+1)
	fmt.Printf("    \tActual  : %v\n", res)
	fmt.Printf("    \tExpected: %v\n", expected)
}

func main() {
	tests := []struct {
		nums1    []int
		nums2    []int
		expected float64
	}{
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{0, 0}, []int{0, 0}, 0},
		{[]int{}, []int{1}, 1},
	}

	passed_all := true
	for i, test := range tests {
		nums1, nums2, expected := test.nums1, test.nums2, test.expected
		res := findMedianSortedArrays(nums1, nums2)
		if res != expected {
			passed_all = false
			print_failed_test(i, res, expected)
		}
	}
	if passed_all {
		fmt.Println("Passed all tests.")
	}
}
