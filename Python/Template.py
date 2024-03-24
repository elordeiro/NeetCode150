from typing import List

class Solution:
    def method(self, x) -> None:
        return


if __name__ == "__main__":
    sol = Solution()
    tests = [
        
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        x, expected = test
        actual = sol.method(x)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------