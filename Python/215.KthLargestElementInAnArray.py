from typing import List
import heapq

class Solution:
    def findKthLargest(self, nums: List[int], k: int) -> int:
        nums = [-1 * num for num in nums]
        heapq.heapify(nums)

        for _ in range(k - 1):
            heapq.heappop(nums)
        
        return -heapq.heappop(nums)


if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([3,2,1,5,6,4], 2, 5),
        ([3,2,3,1,2,4,5,5,6], 4, 4),
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        nums, k, expected = test
        actual = sol.findKthLargest(nums, k)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------