from typing import Optional, List
from TemplateBinaryTree import TreeNode, create_tree
import heapq

null = None

class Solution:
    def kthSmallest(self, root: Optional[TreeNode], k: int) -> int:
        curr = root
        stack = [curr]
        while True:
            while curr:
                stack.append(curr)
                curr = curr.left
            curr = stack.pop()
            k -= 1
            if k == 0:
                return curr.val
            
            curr = curr.right


if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([3,1,4,null,2], 1, 1),
        ([3,1,4,null,2], 2, 2),
        ([5,3,6,2,4,null,null,1], 3, 3),
        ([1,null,2], 2, 2),
        ([2,1,3], 3, 3),
    ]

    passed_all = True
    test_only = 0

    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        nums1, k, expected = test
        root = create_tree(nums1)
        actual = sol.kthSmallest(root, k)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    if passed_all:
        print("All Tests Passed")