from typing import List 
from typing import Optional

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def isSameTree(self, p: Optional[TreeNode], q: Optional[TreeNode]) -> bool:
        def recur(p: Optional[TreeNode], q: Optional[TreeNode]):
            if not p and not q:
                return
            if not p or not q:
                raise
            if p.val != q.val:
                raise
            recur(p.left, q.left)
            recur(q.right, p.right)
        try:
            recur(p,q)
            return True
        except:
            return False

def create_tree(nums: List[int], idx) -> TreeNode:
    if idx >= len(nums):
        return None
    root = TreeNode(val=nums[idx])
    root.left = create_tree(nums, 2 * idx)
    root.right = create_tree(nums, 2 * idx + 1)
    return root

def print_tree(root: Optional[TreeNode]):
    if not root:
        print('[]')
        return
    res = "["
    def recur(root: Optional[TreeNode]):
        nonlocal res
        if not root:
            return
        res += str(root.val) + ', '
        recur(root.left)
        recur(root.right)
    recur(root)
    res = res[:-2]
    res += ']'
    print(res)

def print_test_failed(actual: bool, expected: bool, test_num: int) -> None:
    print(f"Test {test_num} Failed:")
    print(f"\tActual  : {actual}")
    print(f"\tExpected: {expected}")

def compare_tree(tree1: TreeNode, tree2: TreeNode) -> bool:
    if not tree1 and not tree2:
        return True
    if not tree1 or not tree2:
        return False
    if tree1.val != tree2.val:
        return False
    if compare_tree(tree1.left, tree2.left):
        return compare_tree(tree1.right, tree2.right)
    return False

if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([-1,1,2,3], [-1,1,2,3], True),
        ([-1,1,2], [-1,1,None,2], False),
        ([-1,1,2,1], [-1,1,1,2], False),
    ]

    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        nums1, nums2, expected = test
        root1 = create_tree(nums1, 1)
        root2 = create_tree(nums2, 1)
        actual = sol.isSameTree(root1, root2)
        if actual != expected:
            print_test_failed(actual, expected, i)
            passed_all = False
    if passed_all:
        print("All Tests Passed")