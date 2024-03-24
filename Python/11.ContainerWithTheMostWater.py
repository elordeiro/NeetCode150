class Solution:
    def maxArea(self, height: list[int]) -> int:
        max_area = 0
        max_height = max(height)
        i, j = 0, len(height) - 1

        while i < j:
            area = min(height[i], height[j]) * (j - i)
            if area >= max_area:
                max_area = area
            
            if area >= max_height * (j - i):
                break
            
            if height[i] < height[j]:
                i += 1
            else:
                j -= 1
            
        return max_area
 





if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([1,8,6,2,5,4,8,3,7], 49),
        ([1,2,3,4,5,6,7,8], 16),
        ([1,2,4,3], 4),
        ([2,3,4,5,18,17,6], 17)
    ]
    
    passed_all = True
    for i, test in enumerate(tests, 1):
        height, expected = test
        actual = sol.maxArea(height)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

