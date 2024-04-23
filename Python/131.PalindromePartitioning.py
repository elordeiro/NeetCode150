from typing import List

class Solution:
    def isPalindrome(self, s: str, i: int, j: int) -> bool:
        while i < j:
            if s[i] != s[j]:
                return False
            i, j = i + 1, j - 1
        return True


    def partition(self, s: str) -> List[List[str]]:
        n = len(s)
        res =  []
        def backtrack(i, partial):
            if i >= n:
                res.append(partial)
                return
            
            for j in range(i, n):
                if self.isPalindrome(s, i, j):
                    backtrack(j+1, partial + [s[i:j+1]])
        
        backtrack(0, [])
        return res


if __name__ == "__main__":
    sol = Solution()
    tests = [
         ("aab", [["a","a","b"],["aa","b"]]),
         ("a", [["a"]]),
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        s, expected = test
        actual = sol.partition(s)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------