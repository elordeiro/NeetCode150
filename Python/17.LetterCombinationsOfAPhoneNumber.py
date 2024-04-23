from typing import List

class Solution:
    def letterCombinations(self, digits: str) -> List[str]:
        if not digits:
            return []
        res = []
        n = len(digits)
        keypad = { "2": "abc", "3": "def", "4": "ghi", "5": "jkl", "6": "mno", "7": "pqrs", "8": "tuv", "9": "wxyz" }
        def backtrack(i, partial):
            if i == n:
                res.append(partial)
                return
            for char in keypad[digits[i]]:
                partial += char
                backtrack(i+1, partial)
                partial = partial[:-1]
        backtrack(0, "")
        return res



if __name__ == "__main__":
    sol = Solution()
    tests = [
        ("23", ["ad","ae","af","bd","be","bf","cd","ce","cf"]),
        ("", []),
        ("2", ["a","b","c"]),
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        digits, expected = test
        actual = sol.letterCombinations(digits)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------