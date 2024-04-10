from typing import List 
from typing import Optional

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def invertTree(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        if not root:
            return
        root.left, root.right = self.invertTree(root.right), self.invertTree(root.left)
        return root

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

def print_test_failed(actual: Optional[TreeNode], expected: Optional[TreeNode], test_num: int) -> None:
    print(f"Test {test_num} Failed:")
    print(f"\tActual: ", end="")
    print_tree(actual)
    print(f"\tExpected: ", end="")
    print_tree(expected)

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
        ([-1,4,2,7,1,3,6,9], [-1,4,7,2,9,6,3,1])
    ]
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        nums1, nums2 = test
        root = create_tree(nums1, 1)
        actual = sol.invertTree(root)
        expected = create_tree(nums2, 1)
        if not compare_tree(actual, expected):
            print_test_failed(actual, expected, i)
            passed_all = False
    if passed_all:
        print("All Tests Passed")