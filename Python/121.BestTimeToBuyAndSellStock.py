class Solution:
    def maxProfit(self, prices: list[int]) -> int:
        n = len(prices)
        i, j, profit = 0, 0, 0
        
        while j < n:
            if prices[j] < prices[i]:
                i = j
            else:
                profit = max(profit, prices[j] - prices[i])
            j += 1
        
        return profit



if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([7,1,5,3,6,4], 5),
        ([7,2,5,1,7,4], 6),
        ([7,6,4,3,1], 0),
        ([2,1,2,1,0,1,2], 2),
        ([2,4,1,11,7], 10)
    ]
    
    passed_all = True
    for i, test in enumerate(tests, 1):
        prices, expected = test
        actual = sol.maxProfit(prices)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")





# ---------------------------- Other Solutions --------------------------------
def maxProfit(self, prices: list[int]) -> int:
    n = len(prices)
    i, j, profit = 0, 0, 0
    
    while i < n:
        j = i + 1
        while j < n and prices[j] > prices[i]:
            profit = max(profit, prices[j] - prices[i])
            j += 1
        i = j
    return profit 