from typing import List

class Solution:
    def isValid(self, s: str) -> bool:
        mapping = {
            '(': ')',
            '[': ']',
            '{': '}'
        }
        stack = []

        for char in s:
            if char in mapping:
                stack.append(char)
            else:
                if not stack or char != mapping[stack.pop()]:
                    return False

        
        return not stack


if __name__ == "__main__":
    sol = Solution()
    tests = [
        ("()", True),
        ("()[]{}", True),
        ("(]", False),
        ("[", False),
        ("]", False),

    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        s, expected = test
        actual = sol.isValid(s)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")





# ---------------------------- Other Solutions --------------------------------