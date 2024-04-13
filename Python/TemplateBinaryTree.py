from typing import List 
from typing import Optional

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def method(self, head: Optional[TreeNode]) -> Optional[TreeNode]:
        return None

def create_tree(nums: List[int], idx) -> TreeNode:
    if not nums:
        return None
    head = TreeNode(nums.pop(0))
    queue = [head]
    while queue:
        curr = queue.pop(0)
        left = nums.pop(0) if nums else None
        right = nums.pop(0) if nums else None
        curr.left = TreeNode(left)
        curr.right = TreeNode(right)
        if left != None:
            queue.append(curr.left)
        if right != None:
            queue.append(curr.right)
    return head

def print_tree(root: Optional[TreeNode]):
    if not root:
        print('[]')
        return
    res = "["
    stack = []
    queue = [root]
    while queue:
        curr = queue.pop(0)
        if not curr:
            stack.append(None)
            continue
        stack.append(curr.val)
        queue.append(curr.left)
        queue.append(curr.right)
    
    while stack and stack[-1] == None:
        stack.pop()
    
    while stack:
        curr = stack.pop(0)
        if curr == None:
            res += 'null, '
            continue
        res += str(curr) + ', '
    
    res = res[:-2]
    res += ']'
    print(res)

def print_test_failed(actual: Optional[TreeNode], expected: Optional[TreeNode], test_num: int) -> None:
    print(f"Test {test_num} Failed:")
    print(f"\tActual  : ", end="")
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
        ([-1,1,2,3,4,5], [-1,5,4,3,2,1])
    ]

    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        nums1, nums2 = test
        root = create_tree(nums1, 1)
        actual = sol.method(root)
        expected = create_tree(nums2, 1)
        if not compare_tree(actual, expected):
            print_test_failed(actual, expected, i)
            passed_all = False
    if passed_all:
        print("All Tests Passed")