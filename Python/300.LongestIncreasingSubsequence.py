from typing import List

class Solution:
    def lengthOfLIS(self, nums: List[int]) -> int:
        stack = []
        for num in nums:
            n = len(stack)
            l, r = 0, n - 1
            while l <= r:
                mid = (l + r) // 2
                if num > stack[mid]:
                    l = mid + 1
                else:
                    r = mid - 1
            if l == n:
                stack.append(num)
            else:
                stack[l] = num

        return len(stack)

if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([10,9,2,5,3,7,101,18], 4),
        ([0,1,0,3,2,3], 4),
        ([7,7,7,7,7,7,7], 1),
        ([1,3,6,7,9,4,10,5,6], 6)
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        nums, expected = test
        actual = sol.lengthOfLIS(nums)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------