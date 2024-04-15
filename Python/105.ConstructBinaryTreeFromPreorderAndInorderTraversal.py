from typing import Optional, List
from TemplateBinaryTree import TreeNode, create_tree, compare_tree, print_test_failed

null = None

class Solution:
    def buildTree(self, preorder: List[int], inorder: List[int]) -> Optional[TreeNode]:
        if not preorder:
            return None
        val = preorder[0]
        root = TreeNode(val=val)
        idx = inorder.index(val)
        root.left = self.buildTree(preorder[1:idx+1], inorder[:idx])
        root.right = self.buildTree(preorder[idx+1:], inorder[idx+1:])
        return root


if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([3,9,20,15,7], [9,3,15,20,7], [3,9,20,null,null,15,7]),
        ([-1], [-1], [-1]),
    ]

    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        nums1, nums2, expected = test
        actual = sol.buildTree(nums1, nums2)
        expected = create_tree(expected)
        if not compare_tree(actual, expected):
            print_test_failed(actual, expected, i)
            passed_all = False
    if passed_all:
        print("All Tests Passed")