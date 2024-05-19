#include <iostream>
#include <vector>
#include <algorithm>
#include <unordered_map>
#include <unordered_set>
#include <deque>

using namespace std;

class Solution {
public:
    vector<string> findItinerary(vector<vector<string> >& tickets) {
        int flights;
        unordered_map<string, vector<string> > fromSrcToDst;
        vector<string> res;

        sort(tickets.begin(), tickets.end());
        flights = tickets.size();
        
        for (const auto& ticket : tickets) {
            string src = ticket[0];
            string dst = ticket[1];
            if (fromSrcToDst.find(src) == fromSrcToDst.end()) {
                fromSrcToDst[src] = vector<string>();
            }
            fromSrcToDst[src].push_back(dst);
        }
        
        res.emplace_back("JFK");
        function<bool(string)> dfs;
        dfs = [&] (const string& src) {
            if (flights == 0)  {
                return true;
            }
            unordered_set<string> failed;
            for (int i = 0; i < fromSrcToDst[src].size(); i++) {
                auto dst = fromSrcToDst[src][i];
                if (failed.find(dst) != failed.end()) {
                    continue;
                }
                fromSrcToDst[src].erase(fromSrcToDst[src].begin() + i);
                res.push_back(dst);
                flights--;

                if (dfs(dst)) {
                    return true;
                }

                failed.insert(dst);
                fromSrcToDst[src].insert(fromSrcToDst[src].begin() + i, dst);
                flights++;
                res.pop_back();
            }
            return false;
        };
        dfs("JFK");
        return res;
    }
};

struct test {
    vector<vector<string> > tickets;
    vector<string> expected;
};

int main() {
    vector<test> tests {
        {{{"JFK","SFO"},{"JFK","ATL"},{"SFO","ATL"},{"ATL","JFK"},{"ATL","SFO"}}, {"JFK","ATL","JFK","SFO","ATL","SFO"}},
        {{{"MUC","LHR"},{"JFK","MUC"},{"SFO","SJC"},{"LHR","SFO"}}, {"JFK","MUC","LHR","SFO","SJC"}},
        {{{"JFK","SFO"},{"JFK","ATL"},{"SFO","ATL"},{"ATL","JFK"},{"ATL","SFO"}}, {"JFK","ATL","JFK","SFO","ATL","SFO"}},
        {{{"EZE","AXA"},{"TIA","ANU"},{"ANU","JFK"},{"JFK","ANU"},{"ANU","EZE"},{"TIA","ANU"},{"AXA","TIA"},{"TIA","JFK"},{"ANU","TIA"},{"JFK","TIA"}}, {"JFK","ANU","EZE","AXA","TIA","ANU","JFK","TIA","ANU","TIA","JFK"}},
        {{{"AXA","EZE"},{"EZE","AUA"},{"ADL","JFK"},{"ADL","TIA"},{"AUA","AXA"},{"EZE","TIA"},{"EZE","TIA"},{"AXA","EZE"},{"EZE","ADL"},{"ANU","EZE"},{"TIA","EZE"},{"JFK","ADL"},{"AUA","JFK"},{"JFK","EZE"},{"EZE","ANU"},{"ADL","AUA"},{"ANU","AXA"},{"AXA","ADL"},{"AUA","JFK"},{"EZE","ADL"},{"ANU","TIA"},{"AUA","JFK"},{"TIA","JFK"},{"EZE","AUA"},{"AXA","EZE"},{"AUA","ANU"},{"ADL","AXA"},{"EZE","ADL"},{"AUA","ANU"},{"AXA","EZE"},{"TIA","AUA"},{"AXA","EZE"},{"AUA","SYD"},{"ADL","JFK"},{"EZE","AUA"},{"ADL","ANU"},{"AUA","TIA"},{"ADL","EZE"},{"TIA","JFK"},{"AXA","ANU"},{"JFK","AXA"},{"JFK","ADL"},{"ADL","EZE"},{"AXA","TIA"},{"JFK","AUA"},{"ADL","EZE"},{"JFK","ADL"},{"ADL","AXA"},{"TIA","AUA"},{"AXA","JFK"},{"ADL","AUA"},{"TIA","JFK"},{"JFK","ADL"},{"JFK","ADL"},{"ANU","AXA"},{"TIA","AXA"},{"EZE","JFK"},{"EZE","AXA"},{"ADL","TIA"},{"JFK","AUA"},{"TIA","EZE"},{"EZE","ADL"},{"JFK","ANU"},{"TIA","AUA"},{"EZE","ADL"},{"ADL","JFK"},{"ANU","AXA"},{"AUA","AXA"},{"ANU","EZE"},{"ADL","AXA"},{"ANU","AXA"},{"TIA","ADL"},{"JFK","ADL"},{"JFK","TIA"},{"AUA","ADL"},{"AUA","TIA"},{"TIA","JFK"},{"EZE","JFK"},{"AUA","ADL"},{"ADL","AUA"},{"EZE","ANU"},{"ADL","ANU"},{"AUA","AXA"},{"AXA","TIA"},{"AXA","TIA"},{"ADL","AXA"},{"EZE","AXA"},{"AXA","JFK"},{"JFK","AUA"},{"ANU","ADL"},{"AXA","TIA"},{"ANU","AUA"},{"JFK","EZE"},{"AXA","ADL"},{"TIA","EZE"},{"JFK","AXA"},{"AXA","ADL"},{"EZE","AUA"},{"AXA","ANU"},{"ADL","EZE"},{"AUA","EZE"}}, {"JFK", "ADL", "ANU", "ADL", "ANU", "AUA", "ADL", "AUA", "ADL", "AUA", "ANU", "AXA", "ADL", "AUA", "ANU", "AXA", "ADL", "AXA", "ADL", "AXA", "ANU", "AXA", "ANU", "AXA", "EZE", "ADL", "AXA", "EZE", "ADL", "AXA", "EZE", "ADL", "EZE", "ADL", "EZE", "ADL", "EZE", "ANU", "EZE", "ANU", "EZE", "AUA", "AXA", "EZE", "AUA", "AXA", "EZE", "AUA", "AXA", "JFK", "ADL", "EZE", "AUA", "EZE", "AXA", "JFK", "ADL", "JFK", "ADL", "JFK", "ADL", "JFK", "ADL", "TIA", "ADL", "TIA", "AUA", "JFK", "ANU", "TIA", "AUA", "JFK", "AUA", "JFK", "AUA", "TIA", "AUA", "TIA", "AXA", "TIA", "EZE", "AXA", "TIA", "EZE", "JFK", "AXA", "TIA", "EZE", "JFK", "AXA", "TIA", "JFK", "EZE", "TIA", "JFK", "EZE", "TIA", "JFK", "TIA", "JFK", "AUA", "SYD"}},
        {{{"JFK","SFO"},{"JFK","ATL"},{"SFO","JFK"},{"ATL","AAA"},{"AAA","BBB"},{"BBB","ATL"},{"ATL","AAA"},{"AAA","BBB"},{"BBB","ATL"},{"ATL","AAA"},{"AAA","BBB"},{"BBB","ATL"},{"ATL","AAA"},{"AAA","BBB"},{"BBB","ATL"},{"ATL","AAA"},{"AAA","BBB"},{"BBB","ATL"},{"ATL","AAA"},{"AAA","BBB"},{"BBB","ATL"},{"ATL","AAA"},{"AAA","BBB"},{"BBB","ATL"},{"ATL","AAA"},{"AAA","BBB"},{"BBB","ATL"},{"ATL","AAA"},{"AAA","BBB"},{"BBB","ATL"},{"ATL","AAA"},{"AAA","BBB"},{"BBB","ATL"},{"ATL","AAA"},{"AAA","BBB"},{"BBB","ATL"},{"ATL","AAA"},{"AAA","BBB"},{"BBB","ATL"},{"ATL","AAA"},{"AAA","BBB"},{"BBB","ATL"},{"ATL","AAA"},{"AAA","BBB"},{"BBB","ATL"},{"ATL","AAA"},{"AAA","BBB"},{"BBB","ATL"},{"ATL","AAA"},{"AAA","BBB"},{"BBB","ATL"},{"ATL","AAA"},{"AAA","BBB"},{"BBB","ATL"},{"ATL","AAA"},{"AAA","BBB"},{"BBB","ATL"},{"ATL","AAA"},{"AAA","BBB"},{"BBB","ATL"},{"ATL","AAA"},{"AAA","BBB"},{"BBB","ATL"},{"ATL","AAA"},{"AAA","BBB"},{"BBB","ATL"},{"ATL","AAA"},{"AAA","BBB"},{"BBB","ATL"},{"ATL","AAA"},{"AAA","BBB"},{"BBB","ATL"},{"ATL","AAA"},{"AAA","BBB"},{"BBB","ATL"},{"ATL","AAA"},{"AAA","BBB"},{"BBB","ATL"},{"ATL","AAA"},{"AAA","BBB"},{"BBB","ATL"}}, {"JFK","SFO","JFK","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL","AAA","BBB","ATL"}},
    };

    bool failed = false;
    int testOnly = 0;
    int i = 0;
    for (auto test : tests) {
        if (testOnly && testOnly != i + 1) {
            continue;
        }
        i++;
        Solution sol;
        auto actual = sol.findItinerary(test.tickets);
        cout << "Test " << i << " \n\tActual  : ";
        for (int j = 0; j < actual.size(); j++) {
             cout << actual[j] << ", ";
        }
        cout << endl;
        
        cout << "\tExpected: ";
        for (int j = 0; j < test.expected.size(); j++) {
             cout << test.expected[j] << ", ";
        }
        cout << endl;
    }
    if (!failed) {
        cout << "All Tests Passed.\n";
    }
    return 0; 
}