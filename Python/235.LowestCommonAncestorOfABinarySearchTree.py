from typing import List 
from typing import Optional

null = None
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        def recur(root: TreeNode) -> TreeNode:
            if p.val < root.val and q.val < root.val:
                return recur(root.left)
            if p.val > root.val and q.val > root.val:
                return recur(root.right)
            return root
        
        return recur(root)


def create_tree(nums: List[int]) -> TreeNode:
    if not nums:
        return None
    head = TreeNode(nums.pop(0))
    queue = [head]
    while queue:
        curr = queue.pop(0)
        left = nums.pop(0) if nums else None
        right = nums.pop(0) if nums else None
        
        if left != None:
            curr.left = TreeNode(left)
            queue.append(curr.left)
        if right != None:
            curr.right = TreeNode(right)
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
    print(f"\tActual  : {actual.val}")
    print(f"\tExpected: {expected.val}")

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
        ([6,2,8,0,4,7,9,null,null,3,5], 2, 8, 6),
        ([6,2,8,0,4,7,9,null,null,3,5], 2, 4, 2),
        ([2, 1], 2, 1, 2),
    ]

    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        nums1, p, q, expected = test
        root = create_tree(nums1)
        actual = sol.lowestCommonAncestor(root, TreeNode(p), TreeNode(q))
        expected = TreeNode(expected)
        if actual.val != expected.val:
            print_test_failed(actual, expected, i)
            passed_all = False
    if passed_all:
        print("All Tests Passed")