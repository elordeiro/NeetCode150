from typing import List

class Solution:
    def threeSum(self, nums: list[int]) -> list[list[int]]:
        res = []
        n = len(nums)
        i = j = k = 0
        nums = sorted(nums)

        while i < n:
            target = -nums[i]
            j, k = i + 1, n - 1
            while j < k:
                total = nums[j] + nums[k]
                if total < target:
                    j += 1
                elif total > target:
                    k -= 1
                else:
                    res.append([nums[i], nums[j], nums[k]])
                    
                    prev = nums[j]
                    while j < n and prev == nums[j]:
                        j += 1
                    
                    prev = nums[k]
                    while k > j and prev == nums[k]:
                        k -= 1
            
            prev = nums[i]
            while i < n and prev == nums[i]:
                i += 1
        
        return res





if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([-1,0,1,2,-1,-4], [[-1,-1,2],[-1,0,1]]),
        ([0,1,1], []),
        ([0,0,0], [[0,0,0]])
    ]
    
    passed_all = True
    for i, test in enumerate(tests, 1):
        nums, expected = test
        actual = sol.threeSum(nums)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

