from typing import List

class Solution:
    def maxAreaOfIsland(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        def dfs(row: int, col: int) -> int:
            if not(-1 < row < m) or not(-1 < col < n) or grid[row][col] < 1:
                return 0
            grid[row][col] = -1
            return (1 + 
                    dfs(row + 1, col) +
                    dfs(row, col + 1) +
                    dfs(row - 1, col) +
                    dfs(row, col - 1))

        largest = 0
        for row in range(m):
            for col in range(n):
                if grid[row][col] == 1:
                    largest = max(largest, dfs(row, col))
        return largest


if __name__ == "__main__":
    sol = Solution()
    tests = [
        (
            [[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]], 6
        ),(
            [[0,0,0,0,0,0,0,0]], 0
        )
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        grid, expected = test
        actual = sol.maxAreaOfIsland(grid)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------