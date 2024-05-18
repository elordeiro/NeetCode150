from typing import List

class Solution:
    def wordBreak(self, s: str, wordDict: List[str]) -> bool:
        n1 = len(s) + 1
        wordDict = set(wordDict)
        shortest = len(min(wordDict, key=len))
        longest  = len(max(wordDict, key=len))
        dp = [False] * (n1)
        dp[0] = True
        for i in range(shortest, n1):
            for j in range(i-shortest, max(i-longest-1, -1),-1):
                if dp[j] and s[j:i] in wordDict:
                    dp[i] = True
                    break
        return dp[-1]


if __name__ == "__main__":
    sol = Solution()
    tests = [
        ("leetcode", ["leet","code"], True),
        ("applepenapple", ["apple","pen"], True),
        ("catsandog", ["cats","dog","sand","and","cat"], False),
        ("bb", ["a","b","bbb","bbbb"], True),
        ("aaaaaaa", ["aaaa","aaa"], True),
        ("abcd", ["a","abc","b","cd"], True),
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        s, wordDict, expected = test
        actual = sol.wordBreak(s, wordDict)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------