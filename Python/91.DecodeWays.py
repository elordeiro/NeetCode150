class Solution:
    def numDecodings(self, s: str) -> int:
        if not s or s[0] == '0':
            return 0
        
        n = len(s)
        two_back = 1
        one_back = 1

        for i in range(2, n + 1):
            current_digit = s[i - 1]
            prev_digit = s[i - 2]

            temp = 0
            if current_digit != "0":
                temp += one_back

            if prev_digit == "1" or (prev_digit == "2" and current_digit in "0123456"):
                temp += two_back
            
            two_back = one_back
            one_back = temp

        return one_back


if __name__ == "__main__":
    sol = Solution()
    tests = [
        ("12", 2),
        ("226", 3),
        ("06", 0),
        ("10", 1),
        ("27", 1),
        ("2101", 1)
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        s, expected = test
        actual = sol.numDecodings(s)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------