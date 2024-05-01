package main

import (
	"fmt"
	"testing"
)

func Test7ReverseInterger(t *testing.T) {
	tests := []struct {
		x, expected int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
		{900000, 9},
	}
	for i, test := range tests {
		actual := reverse(test.x)
		t.Run("Test7ReverseInterger "+fmt.Sprint(i+1), func(t *testing.T) {
			if actual != test.expected {
				t.Errorf("\n\tActual  : %v\n\tExpected: %v\n", actual, test.expected)
			}
		})
	}
}
