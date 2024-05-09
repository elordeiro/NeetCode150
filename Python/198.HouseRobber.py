from typing import List

class Solution:
    def rob(self, nums: List[int]) -> int:
        if len(nums) == 1:
            return nums[0]
        
        skipped_prev, robbed_prev  = nums[0], nums[1]
        for i in range(2, len(nums)):
            temp = skipped_prev
            skipped_prev = max(robbed_prev, skipped_prev)
            robbed_prev = temp + nums[i]

        return max(skipped_prev, robbed_prev)


if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([1,2,3,1], 4),
        ([2,7,9,3,1], 12),
        ([2,1,1,2], 4),
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