from typing import List
from typing import Optional

# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if not head:
            return
        
        curr = head
        next = curr.next
        curr.next = None
        prev = curr
        curr = next
        while curr:
            next = curr.next
            curr.next = prev
            prev = curr
            curr = next
        
        return prev


def create_list(nums: List[int]) -> ListNode:
    head = ListNode()
    curr = head
    for num in nums:
        curr.next = ListNode(num)
        curr = curr.next
    return head.next

def print_list(head: ListNode) -> None:
    if not head:
        print("]")
        return 
    print(head.val, end=", " if head.next else "")
    print_list(head.next)

def print_test_failed(actual: ListNode, expected: ListNode, test_num: int) -> None:
    print(f"Test {test_num} Failed")
    print(f"\tActual  : [", end="")
    print_list(actual)
    print(f"\tExpected: [", end="")
    print_list(expected)

if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([1,2,3,4,5], [5,4,3,2,1])
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        nums, expected = test
        head = create_list(nums)
        actual = sol.reverseList(head)
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
def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
    if not head:
        return
    stack = []
    curr = head
    while curr:
        stack.append(curr)
        curr = curr.next
    
    head = stack.pop()
    curr = head
    while stack:
        curr.next = stack.pop()
        curr = curr.next
    curr.next = None
    return head