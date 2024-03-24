from collections import defaultdict

class Solution:
    def minWindow(self, s: str, t: str) -> str:
        haveCount = defaultdict(int)
        needCount = defaultdict(int)
        for c in t:
            needCount[c] += 1
        
        minWindow = (0, 0)
        minWindowLen = float('inf')
        l = 0
        have, need = 0, len(needCount.keys())
        for r, char in enumerate(s):
            if char in needCount:
                haveCount[char] += 1
                
                if haveCount[char] == needCount[char]:
                    have += 1
                
                while have == need:
                    if r - l + 1 < minWindowLen:
                        minWindow = (l, r + 1)
                        minWindowLen = r - l + 1
                    
                    if s[l] in needCount:
                        haveCount[s[l]] -= 1
                        if haveCount[s[l]] < needCount[s[l]]:
                            have -= 1
                    l += 1

        return s[minWindow[0]:minWindow[1]]






if __name__ == "__main__":
    sol = Solution()
    tests = [
        ("ADOBECODEBANC", "ABC", "BANC"),
        ("a", "a", "a"),
        ("a", "aa", ""),
        ("a", "b", ""),
        ("ab", "b", "b"),
        ("aa", "aa", "aa"),
        ("bba", "ab", "ba"),
        ("bbaac", "aba", "baa"),
        ("aaaaaaaaaaaabbbbbcdd", "abcdd", "abbbbbcdd")
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        s, t, expected = test
        actual = sol.minWindow(s, t)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")






# ---------------------------- Other Solutions --------------------------------
def minWindow(self, s: str, t: str) -> str:
    l = r = 0
    n, m = len(s), len(t)
    
    tCount = defaultdict(int)
    for c in t:
        tCount[c] += 1
    
    minWindow = ""
    lettersNeeded = m
    while l < n and s[l] not in tCount:
        l += 1

    while r < n:
        if s[r] in tCount:
            tCount[s[r]] -= 1
            if tCount[s[r]] > -1:
                lettersNeeded -= 1
        if lettersNeeded == 0:
            window = s[l : r + 1]
            if not minWindow or len(window) < len(minWindow):
                minWindow = window
            tCount[s[l]] += 1
            if tCount[s[l]] > 0:
                lettersNeeded += 1
            l += 1 
            while l < r and (s[l] not in tCount or tCount[s[l]] < 0):
                if s[l] in tCount:
                    tCount[s[l]] += 1
                l += 1
            if lettersNeeded == 0:
                window = s[l : r + 1]
                if not minWindow or len(window) < len(minWindow):
                    minWindow = window
        r += 1

    return minWindow