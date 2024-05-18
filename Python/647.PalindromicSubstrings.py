class Solution:
    def countSubstrings(self, s: str) -> int:
        count = 0
        n = len(s)
        for i in range(n):
            for j in range(2):
                l, r = i, i + j
                while l > -1 and r < n and s[l] == s[r]:
                    count += 1
                    l -= 1
                    r += 1
        return count


if __name__ == "__main__":
    sol = Solution()
    tests = [
        ("abc", 3),
        ("aaa", 6),
        ("aba", 4),
        ("dnncbwoneinoplypwgbwktmvkoimcooyiwirgbxlcttgteqthcvyoueyftiwgwwxvxvg", 77)
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        s, expected = test
        actual = sol.countSubstrings(s)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------