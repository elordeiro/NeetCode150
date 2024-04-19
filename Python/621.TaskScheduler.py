from typing import List
from collections import deque
import heapq

class Solution:
    def leastInterval(self, tasks: List[str], n: int) -> int:
        mapping =  [0] * 26
        for task in tasks:
            mapping[ord(task) - 65] += 1
        
        heap = []
        for task in mapping:
            if task > 0:
                heapq.heappush(heap, (-task))
        
        t = 0
        q = deque()
        while heap or q:
            
            while q and q[0][1] <= t:
                heapq.heappush(heap, q.popleft()[0])
            
            if heap:
                task_count = heapq.heappop(heap) + 1
                if task_count != 0:
                    q.append((task_count, t + n + 1))
            
            t += 1
        
        return t


if __name__ == "__main__":
    sol = Solution()
    tests = [
        (["A","A","A","B","B","B"], 2, 8),
        (["A","C","A","B","D","B"], 1, 6),
        (["A","A","A", "B","B","B"], 3, 10),
        (["A","A","A","B","B","B"], 50, 104),
        (["A","B","A"], 2, 4),
        (["A","A","A"], 1, 5),
        (["A","A","A","A","A","A","B","C","D","E","F","G"], 1, 12)
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        tasks, n, expected = test
        actual = sol.leastInterval(tasks, n)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------