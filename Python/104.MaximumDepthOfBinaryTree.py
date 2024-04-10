from typing import List 
from typing import Optional

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def maxDepth(self, root: Optional[TreeNode]) -> int:
        if not root:
            return 0
        return 1 + max(self.maxDepth(root.left), self.maxDepth(root.right))

def create_tree(nums: List[int], idx) -> TreeNode:
    if idx >= len(nums) or nums[idx] == None:
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

def print_test_failed(actual: int, expected: int, test_num: int) -> None:
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
        ([-1,3,9,20,None,None,15,7], 3)
    ]

    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        nums1, expected = test
        root = create_tree(nums1, 1)
        actual = sol.maxDepth(root)
        if actual != expected:
            print_test_failed(actual, expected, i)
            passed_all = False
    if passed_all:
        print("All Tests Passed")