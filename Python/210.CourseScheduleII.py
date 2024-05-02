from typing import List

class Solution:
    def findOrder(self, numCourses: int, prerequisites: List[List[int]]) -> List[int]:
        prereqs = [[] for _ in range(numCourses)]
        for course, prereq in prerequisites:
            prereqs[course].append(prereq)
        
        need = set()
        taken = set()
        res = []
        def has_cycle(course: int) -> bool:
            if course in need:
                return True
            if course in taken:
                return False
            
            need.add(course)
            for c in prereqs[course]:
                if has_cycle(c):
                    return True
            need.remove(course)
            taken.add(course)
            res.append(course)
            
            return False

        for course in range(numCourses):
            if has_cycle (course):
                return []
        return res

if __name__ == "__main__":
    sol = Solution()
    tests = [
        (2, [[1,0]], [0,1]),
        (2, [[0,1]], [1,0]),
        (4, [[1,0],[2,0],[3,1],[3,2]], [0,2,1,3]),
        (1, [], [0]),
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        numCourses, prerequisites, expected = test
        actual = sol.findOrder(numCourses, prerequisites)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------