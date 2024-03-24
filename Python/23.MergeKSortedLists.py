from typing import List
from typing import Optional

# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def guard_clauses(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        if (n := len(lists)) > 2:
            return (False, self.mergeKLists([self.mergeKLists(lists[: n // 2]), self.mergeKLists(lists[n // 2 :])]))
        if n == 1:
            return (False, lists[0])
        if n == 0:
            return (False, None)
        
        if not lists[0]:
            return (False, lists[1])
        if not lists[1]:
            return (False, lists[0])
        
        return (True, lists)
    
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        only_two, lists = self.guard_clauses(lists)
        if not only_two:
            return lists
        
        l1 = lists[0]
        l2 = lists[1]
        head = ListNode('-inf')
        head.next = l1 if l1.val < l2.val else l2

        while l1 and l2:
            if l1.val < l2.val:
                while l1.next and l1.next.val < l2.val:
                    l1 = l1.next
                next = l1.next
                l1.next = l2
                l1 = next
            else:
                while l2.next and l2.next.val < l1.val:
                    l2 = l2.next
                next = l2.next
                l2.next = l1
                l2 = next
        
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
        ([[1,4,5],[1,3,4],[2,6]], [1,1,2,3,4,4,5,6]),
        ([], []),
        ([[]], []),
        ([[-2,-1,-1,-1],[]], [-2,-1,-1,-1])
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        nums, expected = test
        lists = []
        while nums:
            lists.append(create_list(nums.pop()))
            
        actual = sol.mergeKLists(lists)
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