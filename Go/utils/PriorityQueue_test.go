package utils

import (
	"container/heap"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	type pair struct {
		fruit    string
		priority int
	}
	tests := []struct {
		foods    []pair
		expected []string
	}{
		{
			[]pair{
				{"banana", 7},
				{"orange", 3},
				{"apple", 1},
				{"grape", 5},
			},
			[]string{"apple", "orange", "grape", "banana"},
		},
		{
			[]pair{
				{"banana", 7},
				{"orange", 3},
				{"apple", 1},
				{"grape", 5},
			},
			[]string{"banana", "grape", "orange", "apple"},
		},
	}

	minHeap := NewPriorityQueue(func(item1, item2 any) bool {
		return item1.(pair).priority < item2.(pair).priority
	})
	for _, pair := range tests[0].foods {
		heap.Push(minHeap, pair)
	}
	t.Run("Testing Min Heap", func(t *testing.T) {
		for _, fruit := range tests[0].expected {
			actual := heap.Pop(minHeap).(pair).fruit
			if fruit != actual {
				t.Errorf("Test Failed\n\tActual  : %v\n\tExpected: %v\n", fruit, actual)
			}
		}
	})

	maxHeap := NewPriorityQueue(func(item1, item2 any) bool {
		return item1.(pair).priority > item2.(pair).priority
	})

	for _, pair := range tests[1].foods {
		heap.Push(maxHeap, pair)
	}
	t.Run("Testing Max Heap", func(t *testing.T) {
		for _, fruit := range tests[1].expected {
			actual := heap.Pop(maxHeap).(pair).fruit
			if fruit != actual {
				t.Errorf("Test Failed\n\tActual  : %v\n\tExpected: %v\n", fruit, actual)
			}
		}
	})

}
