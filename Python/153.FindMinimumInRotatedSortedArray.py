from typing import List

class Solution:
    def findMin(self, nums: List[int]) -> int:
        l, r = 0, len(nums) - 1
        
        if nums[l] <= nums[r]:
            return nums[l]

        mid = (l + r) // 2

        return min([self.findMin(nums[:mid + 1]), self.findMin(nums[mid + 1:])])


if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([3,4,5,1,2], 1),
        ([4,5,6,7,0,1,2], 0),
        ([11,13,15,17], 11),
        ([2,1], 1),
        ([3,1,2], 1)
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        nums, expected = test
        actual = sol.findMin(nums)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------