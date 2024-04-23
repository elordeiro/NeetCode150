from typing import List

class Solution:
    def subsetsWithDup(self, nums: List[int]) -> List[List[int]]:
        nums = sorted(nums)
        res = []
        def backtrack(partial, avail):
            if not avail:
                res.append(partial)
                return
            i = 1
            while i < len(avail) and avail[0] == avail[i]:
                i +=  1
            backtrack(partial, avail[i:])
            backtrack(partial + [avail[0]], avail[1:])
        
        backtrack([], nums)
        return res



if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([1,2,2], [[],[1],[1,2],[1,2,2],[2],[2,2]]),
        ([0], [[],[0]]),
        ([4,4,4,1,4], [[],[1],[1,4],[1,4,4],[1,4,4,4],[1,4,4,4,4],[4],[4,4],[4,4,4],[4,4,4,4]]),
    ]
    
    passed_all = True
    test_only =0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        nums, expected = test
        actual = sol.subsetsWithDup(nums)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------

