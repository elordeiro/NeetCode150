package main

import (
	"reflect"
	"testing"
)

func Test295FindMedianFromDataStream(t *testing.T) {
	tests := []struct {
		ops      []string
		args     [][]int
		expected []float64
	}{
		{
			[]string{"MedianFinder", "addNum", "addNum", "findMedian"},
			[][]int{{}, {0}, {0}, {}},
			[]float64{0, 0, 0, 0},
		},
		{
			[]string{"MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"},
			[][]int{{}, {1}, {2}, {}, {3}, {}},
			[]float64{0, 0, 0, 1.5, 0, 2.0},
		},
		{
			[]string{"MedianFinder", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian"},
			[][]int{{}, {-1}, {}, {-2}, {}, {-3}, {}, {-4}, {}, {-5}, {}},
			[]float64{0, 0, -1, 0, -1.5, 0, -2, 0, -2.5, 0, -3},
		},
		{
			[]string{"MedianFinder", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian"},
			[][]int{{}, {6}, {}, {10}, {}, {2}, {}, {6}, {}, {5}, {}, {0}, {}, {6}, {}, {3}, {}, {1}, {}, {0}, {}, {0}, {}},
			[]float64{0, 0, 6.00000, 0, 8.00000, 0, 6.00000, 0, 6.00000, 0, 6.00000, 0, 5.50000, 0, 6.00000, 0, 5.50000, 0, 5.00000, 0, 4.00000, 0, 3.00000},
		},
		{
			[]string{"MedianFinder", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian"},
			[][]int{{}, {12}, {}, {10}, {}, {13}, {}, {11}, {}, {5}, {}, {15}, {}, {1}, {}, {11}, {}, {6}, {}, {17}, {}, {14}, {}, {8}, {}, {17}, {}, {6}, {}, {4}, {}, {16}, {}, {8}, {}, {10}, {}, {2}, {}, {12}, {}, {0}, {}},
			[]float64{0, 0, 12.00000, 0, 11.00000, 0, 12.00000, 0, 11.50000, 0, 11.00000, 0, 11.50000, 0, 11.00000, 0, 11.00000, 0, 11.00000, 0, 11.00000, 0, 11.00000, 0, 11.00000, 0, 11.00000, 0, 11.00000, 0, 11.00000, 0, 11.00000, 0, 11.00000, 0, 10.50000, 0, 10.00000, 0, 10.50000, 0, 10.00000},
		},
	}
	for _, test := range tests {
		var obj MedianFinder
		actual := []float64{}
		for i, op := range test.ops {
			if op == "MedianFinder" {
				obj = Constructor()
				actual = append(actual, 0)
			} else if op == "addNum" {
				obj.AddNum(test.args[i][0])
				actual = append(actual, 0)
			} else {
				actual = append(actual, obj.FindMedian())
			}
		}
		t.Run("Test295FindMedianFromDataStream", func(t *testing.T) {
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("Test Failed\n\tActual  : %v\n\tExpected: %v", actual, test.expected)
			}
		})
	}
}
