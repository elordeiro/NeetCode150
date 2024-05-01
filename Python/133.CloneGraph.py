from typing import Optional, List
from collections import deque

# Definition for a Node.
class Node:
    def __init__(self, val = 0, neighbors = None):
        self.val = val
        self.neighbors = neighbors if neighbors is not None else []


class Solution:
    def cloneGraph(self, node: Optional['Node']) -> Optional['Node']:
        if not node:
            return None
        mapping = {node : Node(node.val)}
        q = deque([node])
        while q:
            original = q.popleft()
            copy = mapping[original]
            for neighbor in original.neighbors:
                if neighbor not in mapping:
                    mapping[neighbor] = Node(neighbor.val)
                    q.append(neighbor)
                copy.neighbors.append(mapping[neighbor])
        return mapping[node]


def creat_graph(adj_lists: List[List[int]]) -> Node:
    nodes = [None]
    for val in range(1, len(adj_lists) + 1):
        nodes.append(Node(val=val))
    
    for i, adj_list in enumerate(adj_lists, 1):
        curr = nodes[i]
        neighbors = []
        for neighbor in adj_list:
            neighbors.append(nodes[neighbor])
        curr.neighbors = neighbors
    
    return nodes[1]


if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([[2,4],[1,3],[2,4],[1,3]], [[2,4],[1,3],[2,4],[1,3]]),
        ([[]], [[]]),
        ([], []),
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        adj, expected = test
        node = creat_graph(adj)
        actual = sol.cloneGraph(node)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------