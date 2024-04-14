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
    def rightSideView(self, root: Optional[TreeNode]) -> List[int]:
        if not root:
            return []
        res = []
        queue = [root]
        while queue:
            res.append(queue[-1].val)
            for _ in range(len(queue)):
                node = queue.pop(0)
                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)
        return res


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
        ([1,2,3,null,5,null,4], [1,3,4]),
        ([1,null,3], [1,3]),
        ([], []),
    ]

    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        nums1, expected = test
        root = create_tree(nums1)
        actual = sol.rightSideView(root)
        if actual != expected:
            print(f"Test {i} Failed:")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    if passed_all:
        print("All Tests Passed")