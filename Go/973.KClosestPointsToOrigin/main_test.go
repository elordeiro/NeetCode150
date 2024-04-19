package main

import (
	"reflect"
	"testing"
)

func Test973KClosestPointsToOrigin(t *testing.T) {
	tests := []struct {
		points   [][]int
		k        int
		expected [][]int
	}{
		{[][]int{{1, 3}, {-2, 2}}, 1, [][]int{{-2, 2}}},
		{[][]int{{3, 3}, {5, -1}, {-2, 4}}, 2, [][]int{{3, 3}, {-2, 4}}},
	}

	for _, test := range tests {
		points, k, expected := test.points, test.k, test.expected
		actual := kClosest(points, k)
		t.Run("Test973KClosestPointsToOrigin", func(t *testing.T) {
			if len(actual) == 0 && len(expected) != 0 {
				t.Errorf("Test Failed\n\tActual  : %v\n\tExpected: %v\n", actual, expected)
			}
			for _, point := range actual {
				found := false
				for _, elem := range expected {
					if reflect.DeepEqual(point, elem) {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Test Failed\n\tActual  : %v\n\tExpected: %v\n", actual, expected)
				}
			}
		})
	}

}
