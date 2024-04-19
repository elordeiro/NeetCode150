from typing import List
import heapq, math

class Solution:
    def kClosest(self, points: List[List[int]], k: int) -> List[List[int]]:
        if not points:
            return []
        heap = []
        for point in points:
            x, y = point[0], point[1]
            distance = math.sqrt((x ** 2) + (y ** 2))
            heapq.heappush(heap, (-distance, point))
            if len(heap) > k:
                heapq.heappop(heap)
        
        res = []
        for _, point in heap:
            res.append(point)

        return res


if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([[1,3],[-2,2]], 1, [[-2,2]]),
        ([[3,3],[5,-1],[-2,4]], 2, [[3,3],[-2,4]]),
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        points, k, expected = test
        actual = sol.kClosest(points, k)
        if not actual and expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
            continue
        for p in actual:
            if p not in expected:
                print(f"Test {i} Failed")
                print(f"\tActual  : {actual}")
                print(f"\tExpected: {expected}")
                passed_all = False
                break
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------