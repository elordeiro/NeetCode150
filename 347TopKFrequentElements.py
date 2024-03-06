import heapq
from typing import List
from collections import defaultdict

class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        count = {}
        freq = [[] for _ in range((len(nums) + 1))]
        
        for num in nums:
            count[num] = count.get(num, 0) + 1
        
        for key, value in count.items():
            freq[value].append(key)
        
        res = []
        i = len(freq) - 1
        while k > 0:
            j = len(freq[i]) - 1
            while j >= 0:
                res.append(freq[i][j])
                k -= 1
                j -= 1
            i -= 1

        return res


if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([1,1,1,2,2,3], 2, [1,2]),
        ([1], 1, [1])
    ]
    
    passed_all = True
    for i, test in enumerate(tests, 1):
        nums, k, expected = test
        actual = sol.topKFrequent(nums, k)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")





# ---------------------------- Other Solutions --------------------------------
class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        mapping = defaultdict(lambda: 0)
        for num in nums:
            mapping[num] += 1

        heap = []
        for key in mapping.keys():
            heapq.heappush(heap, (-mapping[key], key))
        
        res = []
        for _ in range(k):
            res.append(heapq.heappop(heap)[1])

        return res