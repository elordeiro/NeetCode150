from typing import List
from typing import Optional

# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def reverseKNode(self, head: Optional[ListNode], k: int) -> tuple[Optional[ListNode], Optional[ListNode], Optional[ListNode]]:
        if not head:
            return (None, None, None)
        curr = head
        stack = []
        while curr and k > 0:
            stack.append(curr)
            curr = curr.next
            k -= 1
        next = curr
        if not curr and k > 0:
            return (head, None, None)
        
        curr = stack.pop()
        new_head = curr
        while stack:
            curr.next = stack.pop()
            curr = curr.next
        curr.next = None
        return (new_head, curr, next)
    
    def reverseKGroup(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        res, tail, next = self.reverseKNode(head, k)
        while next:
            new_head, new_tail, next = self.reverseKNode(next, k)
            tail.next = new_head
            tail = new_tail
        return res


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
        ([1,2,3,4,5], 2, [2,1,4,3,5]),
        ([1,2,3,4,5], 3, [3,2,1,4,5])
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        nums, target, expected = test
        head = create_list(nums)
        actual = sol.reverseKGroup(head, target)
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
def reverseKGroup(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
    curr = head
    stack = []
    i = k
    while curr and i > 0:
        stack.append(curr)
        curr = curr.next
        i -= 1
    
    top = stack.pop()
    new_head = top
    while stack:
        top.next = stack.pop()
        top = top.next
    top.next = curr
    old_top = top
    
    while curr:
        stack = []
        i = k
        while curr and i > 0:
            stack.append(curr)
            curr = curr.next
            i -= 1
        
        if len(stack) < k:
            break
        
        top = stack.pop()
        old_top.next = top
        
        while stack:
            top.next = stack.pop()
            top = top.next
        
        old_top = top
        top.next = curr

    return new_head