from typing import List
from collections import defaultdict

class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        mapping = defaultdict(list)
        
        for s in strs:
            mapping[str(sorted(s))].append(s)
        
        return list(mapping.values())


if __name__ == "__main__":
    sol = Solution()
    tests = [
        (["eat","tea","tan","ate","nat","bat"], [["bat"],["nat","tan"],["ate","eat","tea"]]),
        ([""], [[""]]),
        (["a"], [["a"]]),
    ]
    
    passed_all = True
    for i, test in enumerate(tests, 1):
        strs, expected = test
        actual = sol.groupAnagrams(strs)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

