from typing import List 
from typing import Optional

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
     def diameterOfBinaryTree(self, root: Optional[TreeNode]) -> int:
        max_diameter = 0
        def recur(root: Optional[TreeNode]) -> int:
            if not root:
                return -1
            nonlocal max_diameter
            left = recur(root.left)
            right = recur(root.right)
            height = 1 + max(left, right)
            max_diameter = max(max_diameter, left + right + 2)
            return height
        
        recur(root)
        return max_diameter


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
        ([-1,1,2,3,4,5], 3)
    ]

    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        nums1, expected = test
        root = create_tree(nums1, 1)
        actual = sol.diameterOfBinaryTree(root)
        if actual != expected:
            print_test_failed(actual, expected, i)
            passed_all = False
    if passed_all:
        print("All Tests Passed")