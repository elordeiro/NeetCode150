package main

import (
	"math"
	"slices"
	"testing"
)

const null = math.MaxInt

func Test703KthLargestElementInAStream(t *testing.T) {
	tests := []struct {
		inits    [][]int
		ops      []int
		expected []int
	}{
		{
			[][]int{{3}, {4, 5, 8, 2}}, []int{3, 5, 10, 9, 4}, []int{null, 4, 5, 5, 8, 8},
		},
	}
	for _, test := range tests {
		inits, ops, expected := test.inits, test.ops, test.expected
		k, nums := inits[0], inits[1]
		obj := Constructor(k[0], nums)
		actual := []int{null}
		for _, num := range ops {
			actual = append(actual, obj.Add(num))
		}
		t.Run("Test703KthLargestElementInAStream", func(t *testing.T) {
			if !slices.Equal(actual, expected) {
				t.Errorf("\tActual  : %v", actual)
				t.Errorf("\tExpected: %v", expected)
			}
		})
	}
}
