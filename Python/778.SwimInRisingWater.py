from typing import List
from collections import defaultdict
import copy
class Solution:
    def swimInWater(self, grid: List[List[int]]) -> int:
        N = len(grid)
        target_t = 0
        final_r_c = (N-1, N-1)
        visited = set()
        dirs = [(0, 1), (0, -1), (1, 0), (-1, 0)]

        def can_reach_cell(row: int, col: int) -> bool:
            visited.add((row, col))

            for rd, cd in dirs:
                r, c = row + rd, col + cd
                if (-1 < r < N) and (-1 < c < N) and (r, c) not in visited and grid[r][c] <= target_t:
                    if (r, c) == final_r_c:
                        return True
                    if can_reach_cell(r, c):
                        return True
            return False
        
        low = grid[0][0]
        high = ans = N * N - 1

        while low < high:
            mid = (low + high) // 2
            target_t = mid
            visited.clear()
            if can_reach_cell(0, 0):
                ans = mid
                high = mid
            else:
                low = mid + 1
        return ans

if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([[3,2],[0,1]], 3),
        ([[0,2],[1,3]], 3),
        ([[0,1,2,3,4],[24,23,22,21,5],[12,13,14,15,16],[11,17,18,19,20],[10,9,8,7,6]], 16),
        ([[31,28,33,0,8,57,86,99,23,98],[25,90,20,73,34,65,29,9,42,46],[17,84,10,4,40,5,41,21,71,79],[13,70,69,81,63,93,77,1,94,53],[38,87,61,50,92,2,15,95,82,68],[44,72,88,47,27,91,37,48,83,16],[3,30,96,66,7,58,76,54,19,64],[85,45,60,11,51,26,6,22,74,32],[43,12,62,59,89,52,36,97,49,78],[75,24,14,67,56,35,55,39,80,18]], 78)
    ]
    
    passed_all = True
    test_only = 2
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        grid, expected = test
        actual = sol.swimInWater(grid)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------