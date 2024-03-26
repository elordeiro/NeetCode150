from typing import List

class Solution:
    def search(self, nums: List[int], target: int) -> int:
        def binSearch(left, right):
            if nums[left] <= nums[right]:
                while left <= right:
                    mid = (left + right) // 2
                    if target < nums[mid]:
                        right = mid - 1
                    elif nums[mid] < target:
                        left = mid + 1
                    else:
                        return mid
                return -1
            
            mid = (left + right) // 2
            if nums[left] <= nums[mid]:
                if nums[left] <= target <= nums[mid]:
                    return binSearch(left, mid)
                else:
                    return binSearch(mid + 1, right)
            
            if nums[mid] <= target <= nums[right]:
                return binSearch(mid, right)
            else:
                return binSearch(left, mid - 1)
        
        return binSearch(0, len(nums) - 1)


if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([4,5,6,7,0,1,2], 0, 4),
        ([4,5,6,7,0,1,2], 3, -1),
        ([1], 0, -1),
        ([3,1], 0, -1),
        ([5,1,3], 1, 1)
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        nums, target, expected = test
        actual = sol.search(nums, target)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------