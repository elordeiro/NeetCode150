from typing import List

class Solution:
    def maxProduct(self, nums: List[int]) -> int:
        res = nums[0]
        curMax = curMin = 1
        for num in nums:
            temp = curMax
            curMax = max(num, num * curMax, num * curMin)
            curMin = min(num, num * curMin, num * temp)
            res = max(res, curMin, curMax)
        return res

if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([2,3,-2,4], 6),
        ([-2,0,-1], 0),
        ([-2,3,-4], 24),
        ([0, 2], 2),
        ([2,-5,-2,-4,3], 24),
        ([-2], -2),
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        nums, expected = test
        actual = sol.maxProduct(nums)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------