class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        n = len(s)
        l, r = 0, 0
        substring = set()
        longest = 0
        while r < n:
            c = s[r]
            if c not in substring:
                substring.add(c)
            else:
                longest = max(longest, len(substring))
                while s[l] != c:
                    substring.remove(s[l])
                    l += 1
                l += 1
            r += 1

        longest = max(longest, len(substring))
        return longest


if __name__ == "__main__":
    sol = Solution()
    tests = [
        ("abcabcbb", 3),
        ("bbbbb", 1),
        ("pwwkew", 3),
        (" ", 1),
        ("dvdf", 3)
    ]
    
    passed_all = True
    for i, test in enumerate(tests, 1):
        s, expected = test
        actual = sol.lengthOfLongestSubstring(s)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")