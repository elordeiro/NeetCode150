package main

import "testing"

func Test1046LastStoneWeight(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{2, 7, 4, 1, 8, 1}, 1},
		{[]int{1}, 1},
		{[]int{1, 3}, 2},
	}
	for _, test := range tests {
		nums, expected := test.nums, test.expected
		actual := lastStoneWeight(nums)
		t.Run("Test1046LastStoneWeight", func(t *testing.T) {
			if actual != expected {
				t.Errorf("\tActual  : %v", actual)
				t.Errorf("\tExpected: %v", expected)
			}
		})
	}
}
