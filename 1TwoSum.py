from typing import List
from collections import defaultdict

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        mapping = defaultdict(list)
        for i, num in enumerate(nums):
            mapping[num].append(i)
        
        sorted_nums = sorted(nums)
        i, j = 0, len(nums) - 1
        
        while True:
            if (total := (x := sorted_nums[i]) + (y :=sorted_nums[j])) < target:
                i += 1
            elif total > target:
                j -= 1
            else:
                return [mapping[x].pop(), mapping[y].pop()]

if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([2,7,11,15], 9, [0, 1]),
        ([3,2,4], 6, [1,2]),
        ([3,3], 6, [0,1])
    ]
    
    passed_all = True
    for i, test in enumerate(tests, 1):
        nums, target, expected = test
        actual = sol.twoSum(nums, target)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

