# Solution currently only works with positive numbers due to how Python handles negative numbers
# In order to make it work in Python we need to mask off the first 32 digits of the number and check if the result is negative
# Or something like that I'm not sure.

class Solution:
    def getSum(self, a: int, b: int) -> int:
        total, carry = 0, 0
        while True:
            total = a ^ b
            carry = (a & b)  << 1
            a, b = total, carry
            if carry == 0:
                break
        return total


if __name__ == "__main__":
    sol = Solution()
    tests = [
        (-1, 1, 0),
        (1, 2, 3),
        (2, 3, 5),
        (3, 4, 7),
        (20, 30, 50),
    ]
    
    passed_all = True
    test_only = 5
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        a, b, expected = test
        actual = sol.getSum(a, b)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------
