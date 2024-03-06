from typing import List

class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        set_nums = set(nums)
        return len(set_nums) != len(nums)


if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([1,2,3,1], True),
        ([1,2,3,4], False),
        ([1,1,1,3,3,4,3,2,4,2], True)
    ]
    
    passed_all = True
    for i, test in enumerate(tests, 1):
        nums, expected = test
        actual = sol.containsDuplicate(nums)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

