from typing import List

class Solution:
    def minCostConnectPoints(self, points: List[List[int]]) -> int:
        cost = 0
        dist = [float('inf')] * len(points)
        
        def minManhattan(p1: int, p2: int) -> int:
            return min(dist[p2], abs(points[p1][0] - points[p2][0]) + abs(points[p1][1] - points[p2][1]))

        n = len(points)
        for p1 in range(n - 1):
            next = p1 + 1
            for p2 in range(next, n):
                dist[p2] = minManhattan(p1, p2)
                if dist[next] > dist[p2]:
                    dist[next], dist[p2], points[next], points[p2] = dist[p2], dist[next], points[p2], points[next]
            cost += dist[next]

        return cost




if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([[0,0],[2,2],[3,10],[5,2],[7,0]], 20),
        ([[3,12],[-2,5],[-4,1]], 18)
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        points, expected = test
        actual = sol.minCostConnectPoints(points)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------