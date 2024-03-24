class Solution:
    def isPalindrome(self, s: str) -> bool:
        i, j = 0, len(s) - 1
        while i < j:
            while i < j and not s[i].isalnum():
                i += 1
            while i < j and not s[j].isalnum():
                j -= 1

            if s[i].lower() != s[j].lower():
                return False
            
            i += 1
            j -= 1
            
        return True

if __name__ == "__main__":
    sol = Solution()
    tests = [
        ("A man, a plan, a canal: Panama", True),
        ("race a car", False),
        (" ", True)
    ]
    
    passed_all = True
    for i, test in enumerate(tests, 1):
        s, expected = test
        actual = sol.isPalindrome(s)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

