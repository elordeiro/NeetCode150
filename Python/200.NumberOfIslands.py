from typing import List

class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        m, n = len(grid) - 1, len(grid[0]) - 1
        def dfs(row: int, col: int):
            if grid[row][col] == '#' or grid[row][col] == '0':
                return
            grid[row][col] = '#'
            row > 0 and dfs(row - 1, col)
            col > 0 and dfs(row, col - 1)
            row < m and dfs(row + 1, col)
            col < n and dfs(row, col + 1)

        islands = 0
        for row in range(m + 1):
            for col in range(n + 1):
                if grid[row][col] == '1':
                    islands += 1
                    dfs(row, col)
        return islands


if __name__ == "__main__":
    sol = Solution()
    tests = [
        (
            [
                ["1","1","1","1","0"],
                ["1","1","0","1","0"],
                ["1","1","0","0","0"],
                ["0","0","0","0","0"]
            ],
            1
        ),
        (
            [
                ["1","1","0","0","0"],
                ["1","1","0","0","0"],
                ["0","0","1","0","0"],
                ["0","0","0","1","1"]
            ],
            3
        ),
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        grid, expected = test
        actual = sol.numIslands(grid)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------