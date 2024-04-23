package main

import (
	"reflect"
	"testing"
)

func Test131PalindromePartitioning(t *testing.T) {
	tests := []struct {
		s        string
		expected [][]string
	}{
		{"aab", [][]string{{"a", "a", "b"}, {"aa", "b"}}},
		{"a", [][]string{{"a"}}},
	}
	for i, test := range tests {
		actual := partition(test.s)
		t.Run("Test131PalindromePartitioning", func(t *testing.T) {
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("Test %v Failed\n\tActual  : %v\n\tExpected: %v\n", i+1, actual, test.expected)
			}
		})
	}
}
