class WordDictionary:

    def __init__(self):
        self.root = {}

    def addWord(self, word: str) -> None:
        curr = self.root
        for char in word:
            if char not in curr:
                curr[char] = {}
            curr = curr[char]
        curr['*'] = ''

    def dfs(self, i: int, word: str, curr: dict) -> bool:
        if i == len(word):
            return '*' in curr
        if word[i] == '.':
            for char in curr:
                if self.dfs(i+1, word, curr[char]):
                    return True
        if word[i] in curr:
            return self.dfs(i+1, word, curr[word[i]])
        return False

    def search(self, word: str) -> bool:
        return self.dfs(0, word, self.root)


if __name__ == "__main__":
    tests = [
            (
                ["WordDictionary","addWord","addWord","addWord","search","search","search","search"],
                [[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]],
                [None,None,None,None,False,True,True,True],
            ),
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        ops, args, expected = test
        actual = []
        obj = WordDictionary()
        for j, op in enumerate(ops):
            if op == "WordDictionary":
                obj = WordDictionary()
                actual.append(None)
            elif op == "addWord":
                obj.addWord(*args[j])
                actual.append(None)
            elif op == "search":
                actual.append(obj.search(*args[j]))

        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------
