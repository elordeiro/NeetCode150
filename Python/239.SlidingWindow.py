from collections import deque
from typing import List

class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        dq = deque()
        res = []

        for j in range(k):
            while dq and dq[-1] < nums[j]:
                dq.pop()
            dq.append(nums[j])
        
        res.append(dq[0])
        
        i = 0
        for j in range(k, len(nums)):
            if dq[0] == nums[i]:
                dq.popleft()
            while dq and dq[-1] < nums[j]:
                dq.pop()
            dq.append(nums[j])

            res.append(dq[0])
            i += 1
        
        return res
        

if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([1,3,-1,-3,5,3,6,7], 3, [3,3,5,5,6,7]),
        ([1], 1, [1]),
        ([1,-1], 1, [1,-1]),
        ([7,2,4], 2, [7,4])
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        nums, k, expected = test
        actual = sol.maxSlidingWindow(nums, k)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")





# ---------------------------- Other Solutions --------------------------------
def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
    i = 0
    res = []
    for j in range(k, len(nums) + 1):
        largest = float('-inf')
        for num in nums[i:j]:
            largest = max(largest, num)
        res.append(largest)
        i += 1
    return res