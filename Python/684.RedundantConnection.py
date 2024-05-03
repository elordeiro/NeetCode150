from typing import List, Optional
from collections import defaultdict, deque
class Solution:
    def findRedundantConnection(self, edges: List[List[int]]) -> List[int]:
        n = len(edges)
        parent = [i for i in range(n + 1)]
        rank = [1] * (n + 1)
        
        def find(n):
            par = parent[n]
            while par != parent[par]:
                parent[par] = parent[parent[par]]
                par = parent[par]
            return par
        
        def union(src: int, dest: int) -> bool:
            par1, par2, = find(src), find(dest)

            if par1 == par2:
                return False
            
            if rank[par1] > rank[par2]:
                parent[par2] = par1
                rank[par1] += rank[par2]
            else:
                parent[par1] = par2
                rank[par2] += rank[par1]
            
            return True

        for src, dest in edges:
            if not union(src, dest):
                return [src, dest]





if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([[1,2],[1,3],[2,3]], [2,3]),
        ([[1,2],[2,3],[3,4],[1,4],[1,5]], [1,4]),
        ([[1,4],[3,4],[1,3],[1,2],[4,5]], [1,3]),
        ([[2,7],[7,8],[3,6],[2,5],[6,8],[4,8],[2,8],[1,8],[7,10],[3,9]], [2,8]),
        ([[3,4],[1,2],[2,4],[3,5],[2,5]], [2,5]),
        ([[1,5],[2,4],[3,4],[1,3],[3,5]], [3,5]),
    ]
    
    passed_all = True
    test_only = 6
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        edges, expected = test
        actual = sol.findRedundantConnection(edges)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------