from typing import List

class Solution:
    def minCostClimbingStairs(self, cost: List[int]) -> int:
        two_down, one_down = cost[0], cost[1]
        for i in range(2, len(cost)):
            temp = min(one_down, two_down) + cost[i]
            two_down = one_down
            one_down = temp

        return min(one_down, two_down)

if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([10,15,20], 15),
        ([1,100,1,1,1,100,1,1,100,1], 6),
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        cost, expected = test
        actual = sol.minCostClimbingStairs(cost)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------