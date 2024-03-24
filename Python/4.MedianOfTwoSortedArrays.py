from typing import List

class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        half = ((m := len(nums1)) + (n := len(nums2))) // 2
        
        if m > n:
            nums1, nums2 = nums2, nums1
            m, n = n, m
        
        l, r = 0, m - 1
        while True:
            midA = (l + r) // 2
            midB = half - midA - 2

            leftA = nums1[midA] if midA > -1 else float('-inf')
            leftB = nums2[midB] if midB > -1 else float('-inf')
            rightA = nums1[midA + 1] if midA + 1 < m else float('inf')
            rightB = nums2[midB + 1] if midB + 1 < n else float('inf')

            if leftA <= rightB and leftB <= rightA:
                if (m + n) % 2:
                    return min(rightA, rightB)
                return (max(leftA, leftB) + min(rightA, rightB)) / 2
            
            if leftA > rightB:
                r = midA - 1
            else:
                l = midA + 1



if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([1,3], [2], 2.0),
        ([1,2], [3,4], 2.5),
        ([0,0], [0,0], 0),
        ([], [1], 1),
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        nums1, nums2, expected = test
        actual = sol.findMedianSortedArrays(nums1, nums2)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------