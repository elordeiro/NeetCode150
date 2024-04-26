from typing import List

class Solution:

    def create_board(self):
        n = len(self.cols)
        board = [] * n
        for r in range(n):
            row = ""
            for c in range(n):
                if c == self.cols[r]:
                    row += "Q"
                else:
                    row += "."
            board.append(row)
        return board

    def is_consistent(self, row: int, col: int) -> bool:
        for c in self.cols[:row]:
            if c == col:
                return False
        
        neg, pos = row - col, row + col
        for r, c in enumerate(self.cols[:row]):
            if r - c == neg or r + c == pos:
                return False
        
        return True

    def solveNQueens(self, n: int) -> List[List[str]]:
        res = []
        self.cols = [-1] * n
        def backtrack(row):
            if row == n:
                res.append(self.create_board())
                return
            for col in range(n):
                self.cols[row] = col
                if self.is_consistent(row, col):
                    backtrack(row+1)
                self.cols[row] = -1
        backtrack(0)
        return res

if __name__ == "__main__":
    sol = Solution()
    tests = [
        (4, [
                [
                    ".Q..",
                    "...Q",
                    "Q...",
                    "..Q.",
                ],
                [
                    "..Q.",
                    "Q...",
                    "...Q",
                    ".Q..",
                ]
        ]),
        (1, [["Q"]])
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        n, expected = test
        actual = sol.solveNQueens(n)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------