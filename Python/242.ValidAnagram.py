from collections import defaultdict
from collections import Counter
class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        """
        s = sorted(s)
        t = sorted(t)
        return s == t
        """
        if (n := len(s)) != len(t):
            return False
        
        char_freq = defaultdict(lambda: 0)
        for i in range(n):
            char_freq[s[i]] += 1
            char_freq[t[i]] -= 1

        for val in char_freq.values():
            if val:
                return False
        Counter(s)
        return True
    

if __name__ == "__main__":
    sol = Solution()
    tests = [
        ("anagram", "nagaram", True),
        ("rat", "car", False)
    ]
    
    passed_all = True
    for i, test in enumerate(tests, 1):
        s, t, expected = test
        actual = sol.isAnagram(s, t)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")





# ---------------------------- Other Solutions --------------------------------
def isAnagram(self, s: str, t: str) -> bool:
    s = sorted(s)
    t = sorted(t)
    return s == t