class Solution:
    def hammingWeight(self, n: int) -> int:
        total = 0
        while n > 0:
            total += 1 & n
            n >>= 1
        return total


if __name__ == "__main__":
    sol = Solution()
    tests = [
        (
            11, 3
        ),
        (
            128, 1
        ),
        (
            2147483645, 30
        )
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        n, expected = test
        actual = sol.hammingWeight(n)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------