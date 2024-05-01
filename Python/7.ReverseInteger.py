class Solution:
    def reverse(self, x: int) -> int:
        is_negative = False
        if x < 0:
            is_negative = True
            x = -x

        res = 0
        while x > 9:
            res = res * 10 + x % 10
            x //= 10
        
        if res > 214748364 or (res == 214748364 and (is_negative and x > 8 or x > 7)):
            return 0
        
        res = res * 10 + x
        return -res if is_negative else res


if __name__ == "__main__":
    sol = Solution()
    tests = [
        (123, 321),
        (-123, -321),
        (120, 21),
        (900000, 9),
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        x, expected = test
        actual = sol.reverse(x)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------