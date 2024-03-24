from typing import List

class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        i, j = 0, len(numbers) - 1
        while True:
            total = numbers[i] + numbers[j]
            mid = (i + j) // 2
            if total < target:
                i += 1
            elif total > target:
                j -= 1
            else:
                return [i + 1, j + 1]


if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([2,7,11,15], 9, [1,2]),
        ([2,3,4], 6, [1,3]),
        ([-1,0], -1, [1,2]),
        ([3,24,50,79,88,150,345], 200, [3,6])
    ]
    
    passed_all = True
    for i, test in enumerate(tests, 1):
        nums, target, expected = test
        actual = sol.twoSum(nums, target)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

