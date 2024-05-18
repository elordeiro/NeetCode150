from typing import List

class Solution:
    def coinChange(self, coins: List[int], amount: int) -> int:
        memo = [10001] * (amount + 1)
        memo[0] = 0

        for i in range(1, amount + 1):
            memo[i] = min([memo[i - coin] + 1 for coin in coins if i - coin > -1] + [10001])

        return memo[amount] if memo[amount] != 10001 else -1


if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([1,2,5], 11, 3),
        ([2], 3, -1),
        ([1], 0, 0),
        ([2,5,10,1], 27, 4),
        ([186,419,83,408], 6249, 20)
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        coins, amount, expected = test
        actual = sol.coinChange(coins, amount)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------