from typing import List

class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        ops = ["+", "-", "*", "/"]
        stack = []

        for token in tokens:
            if token in ops:
                y = stack.pop()
                x = stack.pop()
                if token == "+":
                    stack.append(x + y)
                elif token == "-":
                    stack.append(x - y)
                elif token == "*":
                    stack.append(x * y)
                else:
                    stack.append(int(x / y))
            else:
                stack.append(int(token))
        
        return stack[0]


if __name__ == "__main__":
    sol = Solution()
    tests = [
        (["2", "1", "+", "3", "*"], 9),
        (["4", "13", "5", "/", "+"], 6),
        (["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"], 22)
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        tokens, expected = test
        actual = sol.evalRPN(tokens)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------