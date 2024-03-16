from typing import List
import math

class Solution:
    def minEatingSpeed(self, piles: List[int], h: int) -> int:
        minK = 1
        maxK = max(piles)
        k = 0
        while minK <= maxK:
            midK = (minK + maxK) // 2
            hours = 0
            for bananas in piles:
                hours += math.ceil(bananas / midK)
                if hours > h:
                    break
            else:
                k = midK
                maxK = midK - 1
                continue
            minK = midK + 1
        return k


if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([3,6,7,11], 8, 4),
        ([30,11,23,4,20], 5, 30),
        ([30,11,23,4,20], 6, 23)
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        piles, h, expected = test
        actual = sol.minEatingSpeed(piles, h)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------