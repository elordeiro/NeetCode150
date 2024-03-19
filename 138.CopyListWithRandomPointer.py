from typing import List
from typing import Optional
from collections import defaultdict

# Definition for singly-linked list.

class Node:
    def __init__(self, x: int, next: 'Node' = None, random: 'Node' = None):
        self.val = int(x)
        self.next = next
        self.random = random


class Solution:
    def copyRandomList(self, head: 'Optional[Node]') -> 'Optional[Node]':
        if not head:
            return None
        mapping = defaultdict(lambda: Node(1))

        curr = head
        while curr:
            node        = mapping[curr]
            node.val    = curr.val
            node.next   = mapping[curr.next] if curr.next else None
            node.random = mapping[curr.random] if curr.random else None
            
            prev = curr
            curr = curr.next
        
        # mapping[prev].next = None
        return mapping[head]



def create_list(nums: List[int]) -> Node:
    head = Node(-1)
    curr = head
    for num in nums:
        curr.next = Node(num)
        curr = curr.next
    return head.next

def print_list(head: Node) -> None:
    if not head:
        print("]")
        return 
    print(head.val, end=", " if head.next else "")
    print_list(head.next)

def print_test_failed(actual: Node, expected: Node, test_num: int) -> None:
    print(f"Test {test_num} Failed")
    print(f"\tActual  : [", end="")
    print_list(actual)
    print(f"\tExpected: [", end="")
    print_list(expected)

if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([7,13,11,10,1], [7,13,11,10,1]),
        ([1,2], [1,2]),
        ([3,3,3], [3,3,3]),
        ([], [])
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        nums, expected = test
        head = create_list(nums)
        actual = sol.copyRandomList(head)
        expected = create_list(expected)
        
        curr_actual = actual
        curr_expected = expected
        while curr_actual and curr_expected:
            if curr_actual.val != curr_expected.val:
                print_test_failed(actual, expected, i)
                passed_all = False
                break
            curr_actual = curr_actual.next
            curr_expected = curr_expected.next
        else:
            if curr_actual or curr_expected:
                print_test_failed(actual, expected, i)
                passed_all = False
            
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------