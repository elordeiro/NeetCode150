from typing import List

class Solution:
    def solve(self, board: List[List[str]]) -> None:
        m, n = len(board) - 1, len(board[0]) - 1

        def dfs(row: int, col: int):
            if board[row][col] == "O":
                board[row][col] = "C"
                row > 0 and dfs(row - 1, col)
                row < m and dfs(row + 1, col)
                col > 0 and dfs(row, col - 1)
                col < n and dfs(row, col + 1)

        for row in range(m + 1):
            if board[row][0] == "O":
                dfs(row, 0)
            if board[row][n] == "O":
                dfs(row, n)
        for col in range(n + 1):
            if board[0][col] == "O":
                dfs(0, col)
            if board[m][col] == "O":
                dfs(m, col)
        
        for row in range(m + 1):
            for col in range(n + 1):
                if board[row][col] == "C":
                    board[row][col] = "O"
                elif board[row][col] == "O":
                    board[row][col] = "X"





if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]], [["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]),
        ([["X"]], [["X"]]),
        ([["O","X","X","O","X"],["X","O","O","X","O"],["X","O","X","O","X"],["O","X","O","O","O"],["X","X","O","X","O"]], [["O","X","X","O","X"],["X","X","X","X","O"],["X","X","X","O","X"],["O","X","O","O","O"],["X","X","O","X","O"]])
    ]
    
    passed_all = True
    test_only = 3
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        board, expected = test
        sol.solve(board)
        if board != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {board}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------