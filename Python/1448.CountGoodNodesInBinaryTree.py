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
    def goodNodes(self, root: TreeNode) -> int:
        if not root:
            return 0
        is_good_sum = 0
        def recur(root: TreeNode, max: int):
            nonlocal is_good_sum
            if root.val >= max:
                is_good_sum += 1
                max = root.val
            if root.left:
                recur(root.left, max)
            if root.right:
                recur(root.right, max)
        recur(root, root.val)
        return is_good_sum

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

if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([3,1,4,3,null,1,5], 4),
        ([3,3,null,4,2], 3),
        ([1], 1),
        ([2,null,4,10,8,null,null,4], 4),
        ([9,null,3,6], 1),
        ([-1,5,-2,4,4,2,-2,null,null,-4,null,-2,3,null,-2,0,null,-1,null,-3,null,-4,-3,3,null,null,null,null,null,null,null,3,-3], 5),
    ]

    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        nums1, expected = test
        root = create_tree(nums1)
        actual = sol.goodNodes(root)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    if passed_all:
        print("All Tests Passed")