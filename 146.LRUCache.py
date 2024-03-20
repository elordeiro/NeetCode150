class Node:
    def __init__(self, key: int, val: int=0, prev: 'Node'=None, next: 'Node'=None) -> None:
        self.key = key
        self.val = val
        self.prev = prev
        self.next = next


class LRUCache:

    def __init__(self, capacity: int):
        self.cap = capacity
        self.len = 0
        self.cache = {}
        
        self.tail = Node(float('-inf'))
        self.head = Node(float('inf'))
        self.tail.next = self.head
        self.head.prev = self.tail

    def get(self, key: int) -> int:
        if key in self.cache:
            node = self.rem_node(key)
            self.add_node(node)
            return node.val
        return - 1

    def put(self, key: int, value: int) -> None:
        if key in self.cache:
            node = self.rem_node(key)
            node.val = value
            self.add_node(node)
            return

        self.add_node(Node(key=key, val=value))

        if self.len < self.cap:
            self.len += 1
            return
        
        self.rem_node(self.tail.next.key)

    def add_node(self, node: 'Node') -> Node:
        node.next = self.head
        node.prev = self.head.prev
        node.prev.next = node
        node.next.prev = node
        self.cache[node.key] = node

    def rem_node(self, key: int) -> Node:
        node = self.cache[key]
        node.prev.next = node.next
        node.next.prev = node.prev
        self.cache.pop(key)
        return node


if __name__ == "__main__":
    tests = [
        ([[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]], [None, None, None, 1, None, -1, None, -1, 3, 4]),
        ([[1],[2,1],[2]], [None,None,1]),
        ([[1],[2,1],[2],[3,2],[2],[3]], [None, None, 1, None, -1, 2]),
        ([[2],[2,1],[2,2],[2],[1,1],[4,1],[2]], [None, None, None, 2, None, None, -1]),
        ([[2],[2],[2,6],[1],[1,5],[1,2],[1],[2]], [None,-1,None,-1,None,None,2,6]),
        ([[10],[10,13],[3,17],[6,11],[10,5],[9,10],[13],[2,19],[2],[3],[5,25],[8],[9,22],[5,5],[1,30],[11],[9,12],[7],[5],[8],[9],[4,30],[9,3],[9],[10],[10],[6,14],[3,1],[3],[10,11],[8],[2,14],[1],[5],[4],[11,4],[12,24],[5,18],[13],[7,23],[8],[12],[3,27],[2,12],[5],[2,9],[13,4],[8,18],[1,7],[6],[9,29],[8,21],[5],[6,30],[1,12],[10],[4,15],[7,22],[11,26],[8,17],[9,29],[5],[3,4],[11,30],[12],[4,29],[3],[9],[6],[3,4],[1],[10],[3,29],[10,28],[1,20],[11,13],[3],[3,12],[3,8],[10,9],[3,26],[8],[7],[5],[13,17],[2,27],[11,15],[12],[9,19],[2,15],[3,16],[1],[12,17],[9,1],[6,19],[4],[5],[5],[8,1],[11,7],[5,2],[9,28],[1],[2,2],[7,4],[4,22],[7,24],[9,26],[13,28],[11,26]], [None,None,None,None,None,None,-1,None,19,17,None,-1,None,None,None,-1,None,-1,5,-1,12,None,None,3,5,5,None,None,1,None,-1,None,30,5,30,None,None,None,-1,None,-1,24,None,None,18,None,None,None,None,-1,None,None,18,None,None,-1,None,None,None,None,None,18,None,None,-1,None,4,29,30,None,12,-1,None,None,None,None,29,None,None,None,None,17,22,18,None,None,None,-1,None,None,None,20,None,None,None,-1,18,18,None,None,None,None,20,None,None,None,None,None,None,None])
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        ops, expected = test
        sol = LRUCache(*ops.pop(0))
        actual = [None]
        while ops:
            op = ops.pop(0)
            if len(op) == 2:
                sol.put(*op)
                actual.append(None)
            else:
                actual.append(sol.get(*op))

        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------