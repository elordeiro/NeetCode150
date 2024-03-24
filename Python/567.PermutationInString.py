from collections import defaultdict
class Solution:
    def checkInclusion(self, s1: str, s2: str) -> bool:
        if (n := len(s1)) > (m := len(s2)): return False
        
        s1_count, s2_count = [0] * 26, [0] * 26
        for i in range(n):
            s1_count[ord(s1[i]) - 97] += 1
            s2_count[ord(s2[i]) - 97] += 1

        matches = 0
        for i in range(26):
            if s1_count[i] == s2_count[i]:
                matches += 1

        i = 0
        for j in range(n, m):
            if all(s1_count[k] == s2_count[k] for k in range(26)): return True

            s2_count[ord(s2[i]) - 97] -= 1
            s2_count[ord(s2[j]) - 97] += 1
            i += 1

        if all(s1_count[k] == s2_count[k] for k in range(26)): return True
        return False


if __name__ == "__main__":
    sol = Solution()
    tests = [
        ("ab", "eidbaooo", True),
        ("ab", "eidboaoo", False),
        ("adc", "dcda", True),
        ("ab", "a", False),
        ("abc", "bbbca", True),
        ("trinitrophenylmethylnitramine", "dinitrophenylhydrazinetrinitrophenylmethylnitramine", True),
    ]
    
    passed_all = True
    for i, test in enumerate(tests, 1):
        s1, s2, expected = test
        actual = sol.checkInclusion(s1, s2)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")



# ---------------------------- Other Solutions --------------------------------
def checkInclusion(self, s1: str, s2: str) -> bool:
        if len(s1) > len(s2):
            return False
        
        s1Count, s2Count = [0] * 26, [0] * 26
        for i in range(len(s1)):
            s1Count[ord(s1[i]) - ord('a')] += 1
            s2Count[ord(s2[i]) - ord('a')] += 1

        matches = 0
        for i in range(26):
            matches += 1 if s1Count[i] == s2Count[i] else 0
        
        l = 0
        for r in range(len(s1), len(s2)):
            if matches == 26: return True

            dummy = []
            for k in range(26):
                char = chr(ord('a') + k)
                dummy.append((char, s1Count[k], s2Count[k]))

            index = ord(s2[r]) - ord('a')
            s2Count[index] += 1
            if s1Count[index] == s2Count[index]:
                matches += 1
            elif s1Count[index] + 1 == s2Count[index]:
                matches -= 1
            
            index = ord(s2[l]) - ord('a')
            s2Count[index] -= 1
            if s1Count[index] == s2Count[index]:
                matches += 1
            elif s1Count[index] - 1 == s2Count[index]:
                matches -= 1

            l += 1
        
        return matches == 26
        