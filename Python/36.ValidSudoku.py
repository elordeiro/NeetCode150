from typing import List
from collections import defaultdict
class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        rows = [set() for _ in range(9)]
        cols = [set() for _ in range(9)]
        sqrs = defaultdict(set)

        for i in range(9):
            for j in range(9):
                cell = board[i][j]
                if cell == '.':
                    continue
                if cell in rows[i] or \
                    cell in cols[j] or \
                    cell in sqrs[(i // 3, j // 3)]:
                    return False
                
                rows[i].add(cell)
                cols[j].add(cell)
                sqrs[(i // 3, j // 3)].add(cell)

        return True



if __name__ == "__main__":
    sol = Solution()
    tests = [
        (
            [
                ["5","3",".",".","7",".",".",".","."]
               ,["6",".",".","1","9","5",".",".","."]
               ,[".","9","8",".",".",".",".","6","."]
               ,["8",".",".",".","6",".",".",".","3"]
               ,["4",".",".","8",".","3",".",".","1"]
               ,["7",".",".",".","2",".",".",".","6"]
               ,[".","6",".",".",".",".","2","8","."]
               ,[".",".",".","4","1","9",".",".","5"]
               ,[".",".",".",".","8",".",".","7","9"]
            ], True
        ),
        (
            [
                ["8","3",".",".","7",".",".",".","."]
               ,["6",".",".","1","9","5",".",".","."]
               ,[".","9","8",".",".",".",".","6","."]
               ,["8",".",".",".","6",".",".",".","3"]
               ,["4",".",".","8",".","3",".",".","1"]
               ,["7",".",".",".","2",".",".",".","6"]
               ,[".","6",".",".",".",".","2","8","."]
               ,[".",".",".","4","1","9",".",".","5"]
               ,[".",".",".",".","8",".",".","7","9"]
            ], False
        ),
        (
            [
                [".",".",".",".","5",".",".","1","."],
                [".","4",".","3",".",".",".",".","."],
                [".",".",".",".",".","3",".",".","1"],
                ["8",".",".",".",".",".",".","2","."],
                [".",".","2",".","7",".",".",".","."],
                [".","1","5",".",".",".",".",".","."],
                [".",".",".",".",".","2",".",".","."],
                [".","2",".","9",".",".",".",".","."],
                [".",".","4",".",".",".",".",".","."]
            ], False
        )
    ]
    
    passed_all = True
    for i, test in enumerate(tests, 1):
        board, expected = test
        actual = sol.isValidSudoku(board)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")


