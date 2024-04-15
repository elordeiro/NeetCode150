from collections import deque
from typing import Optional
from TemplateBinaryTree import TreeNode, create_tree, compare_tree, print_test_failed

null = None

class Codec:
    def serialize(self, root: Optional[TreeNode]) -> str:
        if not root:
            return ""
        stack = []
        dq    = deque([root])
        while dq:
            for _ in range(len(dq)):
                curr = dq.popleft()
                stack.append(curr)
                if curr:
                    dq.append(curr.left)
                    dq.append(curr.right)
        
        while not stack[-1]:
            stack.pop()
        
        res = ""
        for node in stack:
            if node:
                res += str(node.val) + ","
            else:
                res += ","
        return res[:-1]
    
    def deserialize(self, data: str) -> Optional[TreeNode]:
        if not data:
            return None
        nodes = data.split(sep=',')
        n = len(nodes)
        root = TreeNode(val=int(nodes[0]))
        dq = deque([root])
        idx = 1
        while dq:
            curr = dq.popleft()
            left = nodes[idx] if idx < n else None
            idx += 1
            right =  nodes[idx] if idx < n else None
            idx += 1
            if left:
                curr.left = TreeNode(val=int(left))
                dq.append(curr.left)
            if right:
                curr.right = TreeNode(val=int(right))
                dq.append(curr.right)
            
        return root


if __name__ == "__main__":
    ser = Codec()
    deser = Codec()
    
    tests = [
         ([1,2,3,null,null,4,5],  [1,2,3,null,null,4,5]),
         ([1,2,3], [1,2,3]),
         ([1], [1]),
         ([], [])
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        nums1, nums2 = test
        root = create_tree(nums1)
        expected = create_tree(nums2)
        actual = deser.deserialize(ser.serialize(root))
        if not compare_tree(actual, expected):
            print_test_failed(actual, expected, i)
            passed_all = False
    if passed_all:
        print("All Tests Passed")
