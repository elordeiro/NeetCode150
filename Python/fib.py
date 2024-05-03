dp = []

def fib(n: int) -> int:
    if dp[n] != -1:
        return dp[n]
    if n == 1 or n == 2:
        return 1
    dp[n] = fib(n-1) + fib(n-2)
    return dp[n]

n = 1000
dp = [-1] * (n + 1)
print(fib(n))