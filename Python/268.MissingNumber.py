from typing import List

class Solution:
    def missingNumber(self, nums: List[int]) -> int:
        n = len(nums)
        bits = 1 << (n + 1)

        for num in nums:
            bits |= 1 << num 

        i = 0
        while True:
            if not (bits & 1):
                return i
            i += 1
            bits >>= 1


if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([3,0,1], 2),
        ([0,1], 2),
        ([9,6,4,2,3,5,7,0,1], 8),
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        nums, expected = test
        actual = sol.missingNumber(nums)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------