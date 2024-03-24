from typing import List

class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        res = []
        def recur(opened: int, closed: int, partial: str) -> None:
            if opened > n:
                return
            if closed == n:
                res.append(partial)
                return
            if opened == closed:
                return recur(opened + 1, closed, partial + "(")
            
            recur(opened + 1, closed, partial + "(")
            recur(opened, closed + 1, partial + ")")
            
        recur(0, 0, "")
        return res





if __name__ == "__main__":
    sol = Solution()
    tests = [
        (3, ["((()))","(()())","(())()","()(())","()()()"]),
        (1, ["()"])
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        n, expected = test
        actual = sol.generateParenthesis(n)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------