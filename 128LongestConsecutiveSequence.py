from typing import List

class Solution:
    # Fastest solution so far
    def longestConsecutive(self, nums: List[int]) -> int:
        nums = set(nums)
        longest = 0
        while nums:
            curr = nums.pop()
            length = 1
            
            next = curr + 1
            while next in nums:
                nums.remove(next)
                next += 1
                length += 1

            prev = curr - 1
            while prev in nums:
                nums.remove(prev)
                prev -= 1
                length += 1

            longest = max(longest, length)

        return longest


if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([100,4,200,1,3,2], 4),
        ([0,3,7,2,5,8,4,6,0,1], 9),
        ([1,0,-1,2], 4),
        ([0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0], 1)
    ]
    
    passed_all = True
    for i, test in enumerate(tests, 1):
        nums, expected = test
        actual = sol.longestConsecutive(nums)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")





# ---------------------------- Other Solutions --------------------------------
def longestConsecutive(self, nums: List[int]) -> int:
        graph = {}
        for num in  nums:
            graph[num] = True
        
        longest = 0
        visited = {}
        for key in graph.keys():
            curr_len = 0
            i = key
            while i in graph:
                if i in visited:
                    curr_len += visited[i]
                    break
                visited[i] = curr_len
                i += 1
                curr_len += 1

            visited[key] = curr_len
            longest = max(longest, curr_len)
        
        return longest

def longestConsecutive(self, nums: List[int]) -> int:
        set_nums = set(nums)
        longest = 0
        for num in nums:
            if num - 1 not in set_nums:
                length = 0
                while num + length in set_nums:
                    length += 1
                
                longest = max(longest, length)
        
        return longest