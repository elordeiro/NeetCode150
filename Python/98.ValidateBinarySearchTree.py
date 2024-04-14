from typing import List
from typing import Optional

from TemplateBinaryTree import TreeNode, create_tree

null = None

class Solution:
    def isValidBST(self, root: Optional[TreeNode]) -> bool:
        def recur(root: Optional[TreeNode], least: int, greatest: int):
            if not root:
                return True
            if least < root.val < greatest:
                return recur(root.left, least, root.val) and recur(root.right, root.val, greatest)
            return False
        return recur(root, float('-inf'), float('inf'))



if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([2,1,3], True),
        ([5,1,4,null,null,3,6], False),
        ([2,2,2], False),
        ([1,1], False),
        ([0,-1], True),
        ([5,14,null,1], False),
    ]

    passed_all = True
    test_only = 0
    
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        
        nums1, expected = test
        root = create_tree(nums1)
        actual = sol.isValidBST(root)
        
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")