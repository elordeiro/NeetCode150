import math
class Solution:
    def reverseBits(self, n: int) -> int:
        res = 0
        for i in range(32):
            res += n >> i & 1
            res <<= 1
        return res >> 1

if __name__ == "__main__":
    sol = Solution()
    tests = [
        (0b00000010100101000001111010011100, 964176192),
        (0b11111111111111111111111111111101, 3221225471)
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        n, expected = test
        actual = sol.reverseBits(n)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tN       : {n:b}")
            print(f"\tActual  : {actual:b}")
            print(f"\tExpected: {expected:b}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------