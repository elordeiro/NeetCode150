from typing import List

class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        u, d = 0, len(matrix) - 1
        l, r = 0, len(matrix[0]) - 1

        while u <= d:
            m = (u + d) // 2
            if target < matrix[m][0]:
                d = m - 1
            elif matrix[m][r] < target:
                u = m + 1
            else:
                break
        else:
            return False
        
        row = m
        while l <= r:
            m = (l + r) // 2
            if target < matrix[row][m]:
                r = m - 1
            elif matrix[row][m] < target:
                l = m + 1
            else:
                return True
        
        return False




if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([[1,3,5,7],[10,11,16,20],[23,30,34,60]], 3, True),
        ([[1,3,5,7],[10,11,16,20],[23,30,34,60]], 13, False)
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        matrix, target, expected = test
        actual = sol.searchMatrix(matrix, target)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------