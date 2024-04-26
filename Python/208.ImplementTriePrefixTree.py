from collections import defaultdict

class Trie:

    def __init__(self):
        self.root = {}

    def insert(self, word: str) -> None:
        curr = self.root
        for char in word:
            if char not in curr:
                curr[char] = {}
            curr = curr[char]
        curr['*'] = ''
        
    def search(self, word: str) -> bool:
        curr = self.root
        for char in word:
            if char not in curr:
                return False
            curr = curr[char]
        return '*' in curr

    def startsWith(self, prefix: str) -> bool:
        curr = self.root
        for char in prefix:
            if char not in curr:
                return False
            curr = curr[char]
        return True


null = None
if __name__ == "__main__":
    tests = [
        (
            ["Trie", "insert", "search", "search", "startsWith", "insert", "search"],
            [[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]],
            [null, null, True, False, True, null, True],
        ),
        (
            ["Trie", "insert", "search", "search", "startsWith", "insert", "search", "insert", "search"],
            [[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"], ["app"], ["apple"]],
            [null, null, True, False, True, null, True, null, True],
        ),
        (
            ["Trie","insert","insert","insert","insert","insert","insert","search","search","search","search","search","search","search","search","search","startsWith","startsWith","startsWith","startsWith","startsWith","startsWith","startsWith","startsWith","startsWith"],
            [[],["app"],["apple"],["beer"],["add"],["jam"],["rental"],["apps"],["app"],["ad"],["applepie"],["rest"],["jan"],["rent"],["beer"],["jam"],["apps"],["app"],["ad"],["applepie"],["rest"],["jan"],["rent"],["beer"],["jam"]],
            [null,null,null,null,null,null,null,False,True,False,False,False,False,False,True,True,False,True,True,False,False,False,True,True,True]
        )
        
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        ops, args, expected = test
        actual = []
        obj = Trie()
        for i, op in enumerate(ops):
            if op == "Trie":
                actual.append(None)
            elif op == "insert":
                obj.insert(*args[i])
                actual.append(None)
            elif op == "search":
                actual.append(obj.search(*args[i]))
            elif op == "startsWith":
                actual.append(obj.startsWith(*args[i]))

        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------
