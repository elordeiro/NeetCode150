from typing import List
from collections import defaultdict
import heapq

class Solution:
    def findCheapestPrice(self, n: int, flights: List[List[int]], src: int, dst: int, k: int) -> int:
        graph = [[] for _ in range(n)]
        for source, dest, price in flights:
            graph[source].append((dest, price))

        # (stops, price, city)
        minHeap = [(0, 0, src)]
        visited = [float('inf')] * n
        visited[src] = 0
        while minHeap:
            stops, price, city = heapq.heappop(minHeap)
            if stops > k:
                continue
            for dest, newPrice in graph[city]:
                if visited[dest] > price + newPrice:
                    heapq.heappush(minHeap, (stops + 1, price + newPrice, dest))
                    visited[dest] = price + newPrice

        return visited[dst] if visited[dst] != float('inf') else -1 




if __name__ == "__main__":
    sol = Solution()
    tests = [
        (4, [[0,1,100],[1,2,100],[2,0,100],[1,3,600],[2,3,200]], 0, 3, 1, 700),
        (3, [[0,1,100],[1,2,100],[0,2,500]], 0, 2, 1, 200),
        (3, [[0,1,100],[1,2,100],[0,2,500]], 0, 2, 0, 500),
        (3, [[0,1,2],[1,2,1],[2,0,10]], 1, 2, 1, 1),
        (4, [[0,1,1],[0,2,5],[1,2,1],[2,3,1]], 0, 3, 1, 6),
        (11, [[0,3,3],[3,4,3],[4,1,3],[0,5,1],[5,1,100],[0,6,2],[6,1,100],[0,7,1],[7,8,1],[8,9,1],[9,1,1],[1,10,1],[10,2,1],[1,2,100]], 0, 2, 4, 11),
        (8, [[0,1,100],[1,2,200],[2,3,1],[0,4,5],[4,5,4],[5,2,3],[0,6,1],[6,7,1],[7,2,1]], 0, 3, 2, 301)
    ]
    
    passed_all = True
    test_only = 7
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        n, flights, src, dst, k, expected = test
        actual = sol.findCheapestPrice(n, flights, src, dst, k)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------