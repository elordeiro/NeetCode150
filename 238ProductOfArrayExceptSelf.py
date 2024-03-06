from typing import List

class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        num = 1
        res = []
        
        for i in range(len(nums)):
            res.append(num)
            num *= nums[i]

        num = 1
        for i in range(len(nums) - 1, -1, -1):
            res[i] *= num
            num *= nums[i]
        
        return res


if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([1,2,3,4], [24,12,8,6]),
        ([-1,1,0,-3,3], [0,0,9,0,0]),
    ]
    
    passed_all = True
    for i, test in enumerate(tests, 1):
        nums, expected = test
        actual = sol.productExceptSelf(nums)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

