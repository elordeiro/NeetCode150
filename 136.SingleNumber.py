from typing import List

class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        accum = 0
        for num in nums:
            accum ^= num
        return accum

if __name__ == "__main__":
    sol = Solution()
    tests = [
        (
            [2,2,1], 1
        ),
        (
            [4,1,2,1,2], 4
        ),
        (
            [1], 1
        )
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        nums, expected = test
        actual = sol.singleNumber(nums)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------