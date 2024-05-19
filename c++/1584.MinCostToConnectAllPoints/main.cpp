#include <climits>
#include <cmath>
#include <cstdlib>
#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    int minCostConnectPoints(vector<vector<int>>& points) {
        int cost = 0;
        vector<int> dist(points.size(), INT_MAX);
        
        int n = points.size()-1;
        int n1 = points.size();
        for (int p1 = 0; p1 < n; p1++) {
            int next = p1 + 1;
            for (int p2 = next; p2 < n1; p2++) {
                dist[p2] = min(dist[p2], abs(points[p1][0] - points[p2][0]) + abs(points[p1][1] - points[p2][1]));
                if (dist[next] > dist[p2]) {
                    swap(dist[next], dist[p2]);
                    swap(points[next], points[p2]);
                }
            }
            cost += dist[next];
        }
        return cost;
    }
};

struct test {
    vector<vector<int>> points;
    int expected;
};

int main() {
    Solution sol;
    vector<test> tests {
        {{{0,0},{2,2},{3,10},{5,2},{7,0}}, 20},
        {{{3,12},{-2,5},{-4,1}}, 18}
    };

    bool failed = false;
    int testOnly = 0;
    int i = 0;

    for (auto test : tests) {
        if (testOnly && testOnly != i + 1) {
            continue;
        }
        i++;
        int actual = sol.minCostConnectPoints(test.points);
        if (actual != test.expected) {
            failed = true;
            cout << "Test " << i + 1 << " Failed\n\tActual  :" << actual << "\n\tExpected: " << test.expected << endl;
        }
    }
    if (!failed) {
        cout << "All Tests Passed.\n";
    }
    return 0;
}