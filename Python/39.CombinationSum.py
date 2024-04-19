from typing import List

class Solution:
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        res = []
        n = len(candidates) - 1
        def recur(i: int, partialSum: int, partialArr: List[int]):
            if partialSum == target:
                res.append(partialArr)
                return
            if partialSum > target or i > n:
                return
            recur(i, partialSum + candidates[i], partialArr + [candidates[i]])
            recur(i + 1, partialSum, partialArr)

        recur(0, 0, [])
        return res

if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([2,3,6,7], 7, [[2,2,3],[7]]),
        ([2,3,5], 8, [[2,2,2,2],[2,3,3],[3,5]]),
        ([2], 1, []),
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        candidates, target, expected = test
        actual = sol.combinationSum(candidates, target)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------