from typing import List

class Solution:
    def exist(self, board: List[List[str]], word: str) -> bool:
        if any([char not in [char for row in board for char in row] for char in word]):
            return False
        m, n, word_len = len(board), len(board[0]), len(word)
        
        def backtrack(row: int, col: int, i: int) -> bool:
            if i == word_len:
                return True
            if not (-1 < row < m) or not (-1 < col < n) or word[i] != board[row][col] or board[row][col] == '#':
                return False
            
            old = board[row][col]
            board[row][col] = '#'
            res = backtrack(row + 1, col, i + 1) or \
                  backtrack(row - 1, col, i + 1) or \
                  backtrack(row, col + 1, i + 1) or \
                  backtrack(row, col - 1, i + 1)
            board[row][col] = old
            return res

        for i in range(m):
            for j in range(n):
                if backtrack(i, j, 0):
                    return True
        return False



if __name__ == "__main__":
    sol = Solution()
    tests = [
         ([["A","B","C","E"], ["S","F","C","S"],["A","D","E","E"]], "ABCCED", True),
         ([["A","B","C","E"], ["S","F","C","S"],["A","D","E","E"]], "SEE", True),
         ([["A","B","C","E"], ["S","F","C","S"],["A","D","E","E"]], "ABCB", False),
         ([["A","A","A","A","A","A"],["A","A","A","A","A","A"],["A","A","A","A","A","A"],["A","A","A","A","A","A"],["A","A","A","A","A","A"],["A","A","A","A","A","A"]], "AAAAAAAAAAAAAAa", False),
         ([["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], "ABCCED", True),
         ([["a"]], "a", True),
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        board, word, expected = test
        actual = sol.exist(board, word)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------