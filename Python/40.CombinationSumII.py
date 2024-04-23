from typing import List

class Solution:
    def combinationSum2(self, candidates: List[int], target: int) -> List[List[int]]:
        candidates.sort()
        res = []
        def backtrack(idx: int, total: int, partial: List[int]):
            if total == target:
                res.append(partial)
                return

            for i, x in enumerate(candidates[idx:], idx):
                if x + total > target or x > target:
                    break
                if i > idx and x == candidates[i - 1]:
                    continue
    
                backtrack(i + 1, total + x, partial + [x])
        
        backtrack(0, 0, [])
        return res


if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([10,1,2,7,6,1,5], 8, [[1,1,6],[1,2,5],[1,7],[2,6]]),
        ([2,5,2,1,2], 5, [[1,2,2],[5]]),
        ([1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1], 30, [[1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1]]),
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        candidates, target, expected = test
        actual = sol.combinationSum2(candidates, target)
        if any([a not in expected for a in actual]) or any([a not in actual for a in expected]):
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------