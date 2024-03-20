from typing import List
from typing import Optional

# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def hasCycle(self, head: Optional[ListNode]) -> bool:
        if not head:
            return False
        slow = fast = head
        
        while True:
            if slow.next:
                slow = slow.next
            if fast.next and fast.next.next:
                fast = fast.next.next
            else:
                return False
            if slow == fast:
                return True


def create_list(nums: List[int], create_cycle) -> ListNode:
    head = ListNode()
    curr = head
    for num in nums:
        curr.next = ListNode(num)
        curr = curr.next
    
    if create_cycle:
        curr.next = head.next
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
        ([3,2,0,4], True),
        ([1,2], False),
        ([1], False),
        ([], False),
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        nums, expected = test
        head = create_list(nums, expected)
        actual = sol.hasCycle(head)

        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")


# ---------------------------- Other Solutions --------------------------------
def hasCycle(self, head: Optional[ListNode]) -> bool:
        node_set = set()
        curr = head
        while curr:
            if curr in node_set:
                return True
            node_set.add(curr)
            curr = curr.next
        
        return False