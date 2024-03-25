from typing import List

class Solution:
    def carFleet(self, target: int, position: List[int], speed: List[int]) -> int:
        pos_to_speed = {}
        for i, pos in enumerate(position):
            pos_to_speed[pos] = speed[i]
        
        position = sorted(position, reverse=True)
        stack = []
        for pos in position:
            time = (target - pos) / pos_to_speed[pos]
            if not stack or stack[-1] < time:
                stack.append(time)
        
        return len(stack)


if __name__ == "__main__":
    sol = Solution()
    tests = [
        (12, [10,8,0,5,3], [2,4,1,1,3], 3),
        (10, [3], [3], 1),
        (100, [0,2,4], [4,2,1], 1)
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        target, position, speed, expected = test
        actual = sol.carFleet(target, position, speed)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------