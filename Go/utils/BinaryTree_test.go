package utils

import "testing"

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q != nil && p.Val == q.Val {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return false
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if subRoot == nil {
		return true
	}
	if root == nil {
		return false
	}
	if isSameTree(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func TestBinaryTree(t *testing.T) {
	tests := []struct {
		nums1    []int
		nums2    []int
		expected bool
	}{
		{[]int{3, 4, 5, 1, 2}, []int{4, 1, 2}, true},
		{[]int{3, 4, 5, 1, 2, null, null, null, null, 0}, []int{4, 1, 2}, false},
		{[]int{1, 1}, []int{1}, true},
		{[]int{3, 4, 5, 1, 2}, []int{3, 1, 2}, false},
		{[]int{
			-2, -3, -3, -4, -4, -4, -4, -5, -3, null, -5, -5, -5, -5, null, -4, -6, -2, null, -4, -6, -6, -4, -6, -4, -6, -4, null, -3, -5, -5, null, null, null, null, -7, -5, -7, -7, -5, -3, null, -7, -5, null, -5, -7, -5, -5, null, null, null, null, -6, -4, -6, -6, -4, -4, -8, -6, -8, -8, -6, -4, -4, -2, null, null, null, -4, -6, -4, null, -6, -4, -4, -6, -6, -7, -5, -5, null, null, -7, null, null, -3, -3, -5, null, null, null, -7, null, null, -7, -9, -9, -7, null, -5, null, null, -3, null, -1, null, null, -5, -5, -5, -3, -7, -5, null, null, -5, null, null, -5, null, -5, -8, -6, null, -6, -6, null, null, null, -2, null, null, null, null, null, -8, null, -6, -6, -10, null, -10, null, -6, null, null, -4, null, null, null, 0, -6, null, -6, null, -4, null, null, null, -8, -6, -6, -4, -6, null, null, null, null, null, -9, -7, -7, -5, null, null, null, null, null, null, -7, -7, -5, -5, null, null, -11, null, null, null, null, null, -3, -3, null, null, null, null, -7, -5, null, null, null, null, null, -5, null, -5, null, -3, null, null, null, null, -8, -6, -6, -6, null, null, null, null, -8, -6, null, -4, null, null, -12, null, -4, -4, -2, -2, null, -6, null, null, null, -6, null, null, null, null, null, null, -5, null, -7, -5, null, null, -7, -7, null, null, null, null, null, null, null, null, -3, null, null, null, null, -1, null, null, null, null, null, null, -8, -6, null, null, -8, null, null, null, -2, null, null, null, null, null, -7, null, null, null, null, -3, null, null, -4,
		}, []int{
			-3, -4, -4, null, null, -3, null, -2, null, null, -3, -4,
		}, true},
	}

	for _, test := range tests {
		nums1, nums2, expected := test.nums1, test.nums2, test.expected
		root1, root2 := CreatTree(nums1), CreatTree(nums2)
		testname := "Is SubTree"
		t.Run(testname, func(t *testing.T) {
			actual := isSubtree(root1, root2)
			if actual != expected {
				t.Errorf("Actual    : %v", root1.ToString())
				t.Errorf("Expected  : %v", root2.ToString())

			}
		})

	}
}

func TestCompareTree(t *testing.T) {
	tests := []struct {
		nums1    []int
		nums2    []int
		expected bool
	}{
		{[]int{3, 4, 5, 1, 2}, []int{3, 4, 5, 1, 2}, true},
		{[]int{3, 4, 5, 1, 2, null, null, null, null, 0}, []int{3, 4, 5, 1, 2, null, null, null, null, 10}, false},
		{[]int{1, 1}, []int{1, 1}, true},
	}

	for _, test := range tests {
		testname := "Compare Trees"
		nums1, nums2, expected := test.nums1, test.nums2, test.expected
		root1, root2 := CreatTree(nums1), CreatTree(nums2)
		actual1 := CompareTrees(root1, root2)
		actual2 := isSameTree(root1, root2)
		t.Run(testname, func(t *testing.T) {
			if actual1 != actual2 || actual2 != expected {
				t.Errorf("Actual    : %v", root1.ToString())
				t.Errorf("Expected  : %v", root2.ToString())
			}
		})
	}
}
