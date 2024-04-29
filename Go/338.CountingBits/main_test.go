package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test338CountingBits(t *testing.T) {
	tests := []struct {
		n        int
		expected []int
	}{
		{2, []int{0, 1, 1}},
		{5, []int{0, 1, 1, 2, 1, 2}},
	}
	for i, test := range tests {
		actual := countBits(test.n)
		t.Run("Test338CountingBits "+fmt.Sprint(i+1), func(t *testing.T) {
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("\n\tActual  : %v\n\tExpected: %v:", actual, test.expected)
			}
		})
	}
}
