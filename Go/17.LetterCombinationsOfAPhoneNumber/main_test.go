package main

import (
	"reflect"
	"testing"
)

func Test17LetterCombinationsOfAPhoneNumber(t *testing.T) {
	tests := []struct {
		digits   string
		expected []string
	}{
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"", []string{}},
		{"2", []string{"a", "b", "c"}},
	}

	for i, test := range tests {
		actual := letterCombinations(test.digits)
		t.Run("Test17LetterCombinationsOfAPhoneNumber", func(t *testing.T) {
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("Test %v Failed\n\tActual  : %v\n\tExpected: %v", i+1, actual, test.expected)
			}
		})
	}
}
