from typing import List
from collections import defaultdict
import heapq

class Solution:
    def networkDelayTime(self, times: List[List[int]], n: int, k: int) -> int:
        graph = defaultdict(list)
        for src, dst, time in times:
            graph[src].append((dst, time))
        
        dist = [float('inf')] * (n + 1)
        dist[k] = 0

        globalTime = 0
        heap = [(0, k)]
        visitedArr = [False] * (n+1)
        visitedNum = 0

        while heap:
            currTime, node = heapq.heappop(heap)
            if visitedArr[node]:
                continue
            visitedArr[node] = True
            visitedNum += 1
            globalTime = currTime

            for dst, time in graph[node]:
                timeToDst = globalTime + time
                if timeToDst < dist[dst]:
                    dist[dst] = timeToDst
                    heapq.heappush(heap, (timeToDst, dst))

        return globalTime if visitedNum == n else -1


if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([[2,1,1],[2,3,1],[3,4,1]], 4, 2, 2),
        ([[1,2,1]], 2, 1, 1),
        ([[1,2,1]], 2, 2, -1),
        ([[1,2,1],[2,3,2],[1,3,4]], 3, 1, 3),
        ([[1,2,1],[2,3,2],[1,3,2]], 3, 1, 2),
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        times, n, k, expected = test
        actual = sol.networkDelayTime(times, n, k)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------