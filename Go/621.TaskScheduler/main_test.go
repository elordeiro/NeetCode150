package main

import "testing"

func TestTaskScheduler(t *testing.T) {
	tests := []struct {
		tasks    []byte
		n        int
		expected int
	}{
		{[]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2, 8},
		{[]byte{'A', 'C', 'A', 'B', 'D', 'B'}, 1, 6},
		{[]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 3, 10},
		{[]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 50, 104},
		{[]byte{'A', 'B', 'A'}, 2, 4},
		{[]byte{'A', 'A', 'A'}, 1, 5},
		{[]byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}, 1, 12},
	}
	for _, test := range tests {
		actual := leastInterval(test.tasks, test.n)
		t.Run("TestTaskScheduler", func(t *testing.T) {
			if actual != test.expected {
				t.Errorf("Test Failed\n\tActual  : %v\n\tExpected: %v", actual, test.expected)
			}
		})
	}
}
