from typing import List

class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        res = [0] * len(temperatures)
        stack = []
        
        for d, t in enumerate(temperatures):
            while stack and temperatures[stack[-1]] < t:
                pd = stack.pop()
                res[pd] = d - pd
            stack.append(d)

        return res


if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([73,74,75,71,69,72,76,73], [1,1,4,2,1,1,0,0]),
        ([30, 40, 50, 60], [1,1,1,0]),
        ([30,60,90], [1,1,0])
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        temperatures, expected = test
        actual = sol.dailyTemperatures(temperatures)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------