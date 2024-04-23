from typing import List

class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        n = len(nums)
        res = []
        def backtrack(i, partial, curr):
            if len(partial) == n:
                res.append(partial)
                return
            if i >= len(curr):
                return
            backtrack(0, partial + [curr[i]], curr[:i] + curr[i + 1:])
            backtrack(i + 1, partial, curr)
        backtrack(0, [], nums)
        return res




if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([1,2,3], [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]),
        ([0,1], [[0,1],[1,0]]),
        ([1], [[1]]),

    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        nums, expected = test
        actual = sol.permute(nums)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------