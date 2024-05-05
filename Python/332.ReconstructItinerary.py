from typing import List
from collections import defaultdict

class Solution:
    def findItinerary(self, tickets: List[List[str]]) -> List[str]:
        tickets.sort()
        self.flights = len(tickets)
        fromSrcToDst = defaultdict(list)
        for src, dst in tickets:
            fromSrcToDst[src].append(dst)
        
        res = ["JFK"]
        def dfs(src: str) -> bool:
            if self.flights == 0:
                return True
            
            failed = set()
            adj = fromSrcToDst[src]
            for i, dst in enumerate(adj):
                if dst in failed:
                    continue
                
                adj.pop(i)
                res.append(dst)
                self.flights -= 1
                
                if dfs(dst):
                    return True
                
                failed.add(dst)
                adj.insert(i, dst)
                self.flights += 1
                res.pop()
            
            return False
        
        dfs("JFK")
        return res

if __name__ == "__main__":
    sol = Solution()
    tests = [
        ([["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]], ["JFK","MUC","LHR","SFO","SJC"]),
        ([["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]], ["JFK","ATL","JFK","SFO","ATL","SFO"]),
        ([["EZE","AXA"],["TIA","ANU"],["ANU","JFK"],["JFK","ANU"],["ANU","EZE"],["TIA","ANU"],["AXA","TIA"],["TIA","JFK"],["ANU","TIA"],["JFK","TIA"]], ["JFK","ANU","EZE","AXA","TIA","ANU","JFK","TIA","ANU","TIA","JFK"]),
        ([["AXA","EZE"],["EZE","AUA"],["ADL","JFK"],["ADL","TIA"],["AUA","AXA"],["EZE","TIA"],["EZE","TIA"],["AXA","EZE"],["EZE","ADL"],["ANU","EZE"],["TIA","EZE"],["JFK","ADL"],["AUA","JFK"],["JFK","EZE"],["EZE","ANU"],["ADL","AUA"],["ANU","AXA"],["AXA","ADL"],["AUA","JFK"],["EZE","ADL"],["ANU","TIA"],["AUA","JFK"],["TIA","JFK"],["EZE","AUA"],["AXA","EZE"],["AUA","ANU"],["ADL","AXA"],["EZE","ADL"],["AUA","ANU"],["AXA","EZE"],["TIA","AUA"],["AXA","EZE"],["AUA","SYD"],["ADL","JFK"],["EZE","AUA"],["ADL","ANU"],["AUA","TIA"],["ADL","EZE"],["TIA","JFK"],["AXA","ANU"],["JFK","AXA"],["JFK","ADL"],["ADL","EZE"],["AXA","TIA"],["JFK","AUA"],["ADL","EZE"],["JFK","ADL"],["ADL","AXA"],["TIA","AUA"],["AXA","JFK"],["ADL","AUA"],["TIA","JFK"],["JFK","ADL"],["JFK","ADL"],["ANU","AXA"],["TIA","AXA"],["EZE","JFK"],["EZE","AXA"],["ADL","TIA"],["JFK","AUA"],["TIA","EZE"],["EZE","ADL"],["JFK","ANU"],["TIA","AUA"],["EZE","ADL"],["ADL","JFK"],["ANU","AXA"],["AUA","AXA"],["ANU","EZE"],["ADL","AXA"],["ANU","AXA"],["TIA","ADL"],["JFK","ADL"],["JFK","TIA"],["AUA","ADL"],["AUA","TIA"],["TIA","JFK"],["EZE","JFK"],["AUA","ADL"],["ADL","AUA"],["EZE","ANU"],["ADL","ANU"],["AUA","AXA"],["AXA","TIA"],["AXA","TIA"],["ADL","AXA"],["EZE","AXA"],["AXA","JFK"],["JFK","AUA"],["ANU","ADL"],["AXA","TIA"],["ANU","AUA"],["JFK","EZE"],["AXA","ADL"],["TIA","EZE"],["JFK","AXA"],["AXA","ADL"],["EZE","AUA"],["AXA","ANU"],["ADL","EZE"],["AUA","EZE"]], ['JFK', 'ADL', 'ANU', 'ADL', 'ANU', 'AUA', 'ADL', 'AUA', 'ADL', 'AUA', 'ANU', 'AXA', 'ADL', 'AUA', 'ANU', 'AXA', 'ADL', 'AXA', 'ADL', 'AXA', 'ANU', 'AXA', 'ANU', 'AXA', 'EZE', 'ADL', 'AXA', 'EZE', 'ADL', 'AXA', 'EZE', 'ADL', 'EZE', 'ADL', 'EZE', 'ADL', 'EZE', 'ANU', 'EZE', 'ANU', 'EZE', 'AUA', 'AXA', 'EZE', 'AUA', 'AXA', 'EZE', 'AUA', 'AXA', 'JFK', 'ADL', 'EZE', 'AUA', 'EZE', 'AXA', 'JFK', 'ADL', 'JFK', 'ADL', 'JFK', 'ADL', 'JFK', 'ADL', 'TIA', 'ADL', 'TIA', 'AUA', 'JFK', 'ANU', 'TIA', 'AUA', 'JFK', 'AUA', 'JFK', 'AUA', 'TIA', 'AUA', 'TIA', 'AXA', 'TIA', 'EZE', 'AXA', 'TIA', 'EZE', 'JFK', 'AXA', 'TIA', 'EZE', 'JFK', 'AXA', 'TIA', 'JFK', 'EZE', 'TIA', 'JFK', 'EZE', 'TIA', 'JFK', 'TIA', 'JFK', 'AUA', 'SYD']),
        ([["JFK","SFO"],["JFK","ATL"],["SFO","JFK"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"]], ["JFK","SFO","JFK","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL"])
    ]
    
    passed_all = True
    test_only = 0
    for i, test in enumerate(tests, 1):
        if test_only and test_only != i:
            continue
        tickets, expected = test
        actual = sol.findItinerary(tickets)
        if actual != expected:
            print(f"Test {i} Failed")
            print(f"\tActual  : {actual}")
            print(f"\tExpected: {expected}")
            passed_all = False
    
    if passed_all:
        print("All Tests Passed")

# ---------------------------- Other Solutions --------------------------------