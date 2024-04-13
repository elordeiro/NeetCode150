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
    
    def isSubtree(self, root: Optional[TreeNode], subRoot: Optional[TreeNode]) -> bool:
        if not subRoot:
            return True
        if not root:
            return False
        if self.isSameTree(root, subRoot):
            return True
        return self.isSubtree(root.left, subRoot) or self.isSubtree(root.right, subRoot)

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
        ([3,4,5,1,2], [4,1,2], True),
        ([3,4,5,1,2,None,None,None,None,0], [4,1,2], False),
        ([1,1], [1], True),
        ([3,4,5,1,None,2], [3,1,2], False),
        (
            [
                -2,-3,-3,-4,-4,-4,-4,-5,-3,None,-5,-5,-5,-5,None,-4,-6,-2,None,-4,-6,-6,-4,-6,-4,-6,-4,None,-3,-5,-5,None,None,None,None,-7,-5,-7,-7,-5,-3,None,-7,-5,None,-5,-7,-5,-5,None,None,None,None,-6,-4,-6,-6,-4,-4,-8,-6,-8,-8,-6,-4,-4,-2,None,None,None,-4,-6,-4,None,-6,-4,-4,-6,-6,-7,-5,-5,None,None,-7,None,None,-3,-3,-5,None,None,None,-7,None,None,-7,-9,-9,-7,None,-5,None,None,-3,None,-1,None,None,-5,-5,-5,-3,-7,-5,None,None,-5,None,None,-5,None,-5,-8,-6,None,-6,-6,None,None,None,-2,None,None,None,None,None,-8,None,-6,-6,-10,None,-10,None,-6,None,None,-4,None,None,None,0,-6,None,-6,None,-4,None,None,None,-8,-6,-6,-4,-6,None,None,None,None,None,-9,-7,-7,-5,None,None,None,None,None,None,-7,-7,-5,-5,None,None,-11,None,None,None,None,None,-3,-3,None,None,None,None,-7,-5,None,None,None,None,None,-5,None,-5,None,-3,None,None,None,None,-8,-6,-6,-6,None,None,None,None,-8,-6,None,-4,None,None,-12,None,-4,-4,-2,-2,None,-6,None,None,None,-6,None,None,None,None,None,None,-5,None,-7,-5,None,None,-7,-7,None,None,None,None,None,None,None,None,-3,None,None,None,None,-1,None,None,None,None,None,None,-8,-6,None,None,-8,None,None,None,-2,None,None,None,None,None,-7,None,None,None,None,-3,None,None,-4
            ], 
            [
                -3,-4,-4,None,None,-3,None,-2,None,None,-3,-4
            ],
            True
        ),
    ]

    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        nums1, nums2, expected = test
        root1 = create_tree(nums1, 1)
        root2 = create_tree(nums2, 1)
        actual = sol.isSubtree(root1, root2)
        if actual != expected:
            print_test_failed(actual, expected, i)
            passed_all = False
    if passed_all:
        print("All Tests Passed")
