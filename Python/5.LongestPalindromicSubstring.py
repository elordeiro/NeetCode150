class Solution:
    def longestPalindrome(self, s: str) -> str:
        
        n = len(s)
        self.res = ""

        def expand(l: int, r: int):
            while l > -1 and r < n and s[l] == s[r]:
                if r - l + 1 > len(self.res):
                    self.res = s[l:r+1]
                l -= 1
                r += 1
    
        for i in range(n):
            expand(i, i)
            expand(i, i+1)
        return self.res

if __name__ == "__main__":
    sol = Solution()
    tests = [
        ("babad", "bab"),
        ("cbbd", "bb"),
        ("a", "a"),
        ("babaddtattarrattatddetartrateedredividerb", "ddtattarrattatdd"),
        ("bbbb", "bbbb"),
        ("slvafhpfjpbqbpcuwxuexavyrtymfydcnvvbvdoitsvumbsvoayefsnusoqmlvatmfzgwlhxtkhdnlmqmyjztlytoxontggyytcezredlrrimcbkyzkrdeshpyyuolsasyyvxfjyjzqksyxtlenaujqcogpqmrbwqbiaweacvkcdxyecairvvhngzdaujypapbhctaoxnjmwhqdzsvpyixyrozyaldmcyizilrmmmvnjbyhlwvpqhnnbausoyoglvogmkrkzppvexiovlxtmustooahwviluumftwnzfbxxrvijjyfybvfnwpjjgdudnyjwoxavlyiarjydlkywmgjqeelrohrqjeflmdyzkqnbqnpaewjdfmdyoazlznzthiuorocncwjrocfpzvkcmxdopisxtatzcpquxyxrdptgxlhlrnwgvee", "")
    ]
    
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        s, expected = test
        actual = sol.longestPalindrome(s)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------