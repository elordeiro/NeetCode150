import heapq

class MedianFinder:

    def __init__(self):
        self.leftOfMedian = [] # Is a Max Heap
        self.riteOfMedian = [] # Is a Min Heap
        self.isEven = True

    def addNum(self, num: int) -> None:
        if self.isEven:
            heapq.heappush(self.riteOfMedian,-heapq.heappushpop(self.leftOfMedian, -num))
        else:
            heapq.heappush(self.leftOfMedian, -heapq.heappushpop(self.riteOfMedian, num))
        
        self.isEven = not self.isEven

    def findMedian(self) -> float:
        if self.isEven:
            return (-self.leftOfMedian[0] + self.riteOfMedian[0]) / 2
        return self.riteOfMedian[0]



null = None
if __name__ == "__main__":
    tests = [
        (
            ["MedianFinder","addNum","addNum","findMedian"],
            [[],[0],[0],[]],
            [null, null, null, 0]
        ),
        (
            ["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"],
            [[], [1], [2], [], [3], []],
            [null, null, null, 1.5, null, 2.0],
        ),
        (
            ["MedianFinder","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian"],
            [[],[-1],[],[-2],[],[-3],[],[-4],[],[-5],[]],
            [null, null, -1, null, -1.5, null, -2, null, -2.5, null, -3]
        ),
        (
            ["MedianFinder","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian"],
            [[],[6],[],[10],[],[2],[],[6],[],[5],[],[0],[],[6],[],[3],[],[1],[],[0],[],[0],[]],
            [null,null,6.00000,null,8.00000,null,6.00000,null,6.00000,null,6.00000,null,5.50000,null,6.00000,null,5.50000,null,5.00000,null,4.00000,null,3.00000]
        ),
        (
            ["MedianFinder","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian"],
            [[],[12],[],[10],[],[13],[],[11],[],[5],[],[15],[],[1],[],[11],[],[6],[],[17],[],[14],[],[8],[],[17],[],[6],[],[4],[],[16],[],[8],[],[10],[],[2],[],[12],[],[0],[]],
            [null,null,12.00000,null,11.00000,null,12.00000,null,11.50000,null,11.00000,null,11.50000,null,11.00000,null,11.00000,null,11.00000,null,11.00000,null,11.00000,null,11.00000,null,11.00000,null,11.00000,null,11.00000,null,11.00000,null,11.00000,null,10.50000,null,10.00000,null,10.50000,null,10.00000]
        )
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        ops, args, expected = test
        obj = None
        actual = []
        for j, op in enumerate(ops):
            if op == "MedianFinder":
                obj = MedianFinder()
                actual.append(None)
            elif op == "addNum":
                actual.append(obj.addNum(args[j][0]))
            elif op == "findMedian":
                actual.append(float(obj.findMedian()))

        if actual != expected:
            print(f"Test {i} Failed")
            print(f"Actual  : {actual}")
            print(f"Expected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------