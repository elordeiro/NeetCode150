from collections import defaultdict

class Solution:
    def characterReplacement(self, s: str, k: int) -> int:
        l = 0
        mostFreq = 0
        mapping  = defaultdict(int)
        
        for r in range(len(s)):
            mapping[s[r]] += 1
            mostFreq = max(mostFreq, mapping[s[r]])

            if (r - l + 1 - mostFreq) > k:
                mapping[s[l]] -= 1
                l += 1
    
        return r - l + 1


if __name__ == "__main__":
    sol = Solution()
    tests = [
        ("ABAB", 2, 4),
        ("AABABBA", 1, 4),
        ("AABALMNOPQRS", 1, 4)
    ]
    
    passed_all = True
    for i, test in enumerate(tests, 1):
        s, k, expected = test
        actual = sol.characterReplacement(s, k)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")