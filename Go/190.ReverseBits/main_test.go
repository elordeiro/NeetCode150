package main

import (
	"fmt"
	"testing"
)

func Test190ReverseBits(t *testing.T) {
	tests := []struct {
		n        uint32
		expected uint32
	}{
		{0b00000010100101000001111010011100, 964176192},
		{0b11111111111111111111111111111101, 3221225471},
	}
	for i, test := range tests {
		actual := reverseBits(test.n)
		t.Run("Test190ReverseBits "+fmt.Sprint(i+1), func(t *testing.T) {
			if actual != test.expected {
				t.Errorf("\n\tActual  : %b\n\tExpected: %b", actual, test.expected)
			}
		})
	}
}
