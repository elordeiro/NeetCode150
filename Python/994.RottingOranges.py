from typing import List
from collections import deque

class Solution:
    def orangesRotting(self, grid: List[List[int]]) -> int:
        m, n = len(grid) - 1, len(grid[0]) - 1
        queue = deque()
        to_rot = set()

        for row in range(m + 1):
            for col in range(n + 1):
                if (cell := grid[row][col]) == 1:
                    to_rot.add((row, col))
                elif cell == 2:
                    to_rot.add((row, col))
                    queue.append((row, col))
        
        if not to_rot:
            return 0

        time = -1
        while queue:
            time += 1
            for _ in range(len(queue)):
                row, col = queue.popleft()
                to_rot.remove((row, col))
                if row > 0 and grid[row - 1][col] == 1: # up
                    grid[row - 1][col] = 2
                    queue.append((row - 1, col))
                if row < m and grid[row + 1][col] == 1: # down
                    grid[row + 1][col] = 2
                    queue.append((row + 1, col))
                if col > 0 and grid[row][col - 1] == 1: # left
                    grid[row][col - 1] = 2
                    queue.append((row, col - 1))
                if col < n and grid[row][col + 1] == 1: # right
                    grid[row][col + 1] = 2
                    queue.append((row, col + 1))
        
        if to_rot:
            return -1
        return time


if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([[2,1,1],[1,1,0],[0,1,1]], 4),
        ([[2,1,1],[0,1,1],[1,0,1]], -1),
        ([[0,2]], 0),
        ([[0]], 0),
        ([[1]], -1),
        ([[2,2],[1,1],[0,0],[2,0]], 1)
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        grid, expected = test
        actual = sol.orangesRotting(grid)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------