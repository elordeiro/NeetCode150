from typing import Optional, List
from TemplateBinaryTree import TreeNode, create_tree, compare_tree, print_test_failed

null = None

class Solution:
    def maxPathSum(self, root: Optional[TreeNode]) -> int:
        maxSum = float('-inf')
        def traverse(root: Optional[TreeNode]) -> int:
            if not root:
                return 0
            nonlocal maxSum
            left, right = traverse(root.left), traverse(root.right)
            left, right = max(left, 0), max(right, 0)
            maxSum = max(maxSum, left + root.val + right)
            return root.val + max(left, right)
        traverse(root)
        return maxSum

if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([1,2,3], 6),
        ([-10,9,20,null,null,15,7], 42),
        ([-3], -3),
        ([2,-1], 2),
        ([1,2], 3),
        ([1,-2,-3,1,3,-2,null,-1], 3),
        ([5,4,8,11,null,13,4,7,2,null,null,null,1], 48)
    ]

    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        nums1, expected = test
        actual = sol.maxPathSum(create_tree(nums1))
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    if passed_all:
        print("All Tests Passed")