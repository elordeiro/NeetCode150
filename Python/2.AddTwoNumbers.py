from typing import List
from typing import Optional

# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        carry = 0
        head = ListNode()
        curr = head
        while l1 and l2:
            x, y   = l1.val, l2.val
            l1, l2 = l1.next, l2.next
            d, m = divmod(x + y + carry, 10)
            curr.next = ListNode(m)
            curr = curr.next
            carry = d
        
        while l1:
            x  = l1.val
            l1 = l1.next
            d, m = divmod(x + carry, 10)
            curr.next = ListNode(m)
            curr = curr.next
            carry = d
            
        while l2:
            y  = l2.val
            l2 = l2.next
            d, m = divmod(y + carry, 10)
            curr.next = ListNode(m)
            curr = curr.next
            carry = d
            
        if carry:
            curr.next = ListNode(carry)
        
        return head.next


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
        ([2,4,3], [5,6,4], [7,0,8]),
        ([0], [0], [0]),
        ([9,9,9,9,9,9,9], [9,9,9,9], [8,9,9,9,0,0,0,1]),
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        l1, l2, expected = test
        l1 = create_list(l1)
        l2 = create_list(l2)
        actual = sol.addTwoNumbers(l1, l2)
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