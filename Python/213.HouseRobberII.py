from typing import List

class Solution:
    def rob(self, nums: List[int]) -> int:
        if len(nums) == 1:
            return nums[0]
        
        skipped_prev1, robbed_prev1  = nums[0], nums[1]
        skipped_prev2, robbed_prev2  = 0, nums[1]
        for i in range(2, len(nums) - 1):
            temp1, temp2 = skipped_prev1, skipped_prev2
            skipped_prev1, skipped_prev2 = max(robbed_prev1, skipped_prev1), max(robbed_prev2, skipped_prev2)
            robbed_prev1, robbed_prev2 = temp1 + nums[i], temp2 + nums[i]
        
        temp = skipped_prev2
        skipped_prev2 = max(robbed_prev2, skipped_prev2)
        robbed_prev2 = temp + nums[-1]
        
        return max([skipped_prev1, robbed_prev1, skipped_prev2, robbed_prev2])

if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([2,3,2], 3),
        ([1,2,3,1], 4),
        ([1,2,3], 3),
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        nums, expected = test
        actual = sol.rob(nums)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------