package utils

import "testing"

func TestLinkedList(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		list1, list2 := CreateList(test.nums1), CreateList(test.nums2)
		testname := "Compare Lists"
		t.Run(testname, func(t *testing.T) {
			if !CompareLists(list1, list2) {
				t.Errorf("Actual   : %v", list1.ToString())
				t.Errorf("Expected : %v", list2.ToString())
			}
		})
	}
}
