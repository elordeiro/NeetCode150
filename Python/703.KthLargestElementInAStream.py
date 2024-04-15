from typing import List
import heapq

class KthLargest:

    def __init__(self, k: int, nums: List[int]):
        self.k = k
        self.heap = []
        for num in nums:
            heapq.heappush(self.heap, num)

    def add(self, val: int) -> int:
        heapq.heappush(self.heap, val)
        while len(self.heap) > self.k:
            heapq.heappop(self.heap)
        return self.heap[0]

null = None
if __name__ == "__main__":
    tests = [
        ([[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]], [null, 4, 5, 5, 8, 8]),
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        ops, expected = test
        first = ops.pop(0)
        k, nums = first[0], first[1]
        obj = KthLargest(k, nums)
        actual = [None]
        for num in ops:
            actual.append(obj.add(num[0]))
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------


# type KthLargest struct {
# 	k    int
# 	heap *IntHeap
# }

# func Constructor(k int, nums []int) KthLargest {
# 	h := &IntHeap{}
# 	heap.Init(h)
# 	for _, num := range nums {
# 		heap.Push(h, num)
# 		if h.Len() > k {
# 			heap.Pop(h)
# 		}
# 	}
# 	return KthLargest{k: k, heap: h}
# }

# func (this *KthLargest) Add(val int) int {
# 	heap.Push(this.heap, val)
# 	if this.heap.Len() > this.k {
# 		heap.Pop(this.heap)
# 	}
# 	return (*this.heap)[0]
# }

# type IntHeap []int

# func (h IntHeap) Len() int {
# 	return len(h)
# }

# func (h IntHeap) Less(i, j int) bool {
# 	return h[i] < h[j]
# }

# func (h IntHeap) Swap(i, j int) {
# 	h[i], h[j] = h[j], h[i]
# }

# func (h *IntHeap) Push(x interface{}) {
# 	*h = append(*h, x.(int))
# }

# func (h *IntHeap) Pop() interface{} {
# 	old := *h
# 	n := len(old)
# 	x := old[n-1]
# 	*h = old[:n-1]
# 	return x
# }

# func main() {
# 	nums := []int{1, 2, 3, 4, 5, 6}

# }