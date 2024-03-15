from typing import List

class Solution:
    def largestRectangleArea(self, heights: List[int]) -> int:
        stack = []
        largest = 0
        for j, h in enumerate(heights):
            i = j
            while stack and stack[-1][1] > h:
                i, height = stack.pop()
                largest = max(largest, height * (j - i))

            stack.append((i, h))
        
        j += 1
        while stack:
            i, height = stack.pop()
            largest = max(largest, height * (j - i))
        
        return largest


if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([2,1,5,6,2,3], 10),
        ([2,4], 4)
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        heights, expected = test
        actual = sol.largestRectangleArea(heights)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------