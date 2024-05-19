#include <climits>
#include <vector>
#include <iostream>
#include <queue>
#include <unordered_map>

using namespace std;

struct Edge {
    int dst, time;
};

class Solution {
public:
    int networkDelayTime(vector<vector<int>>& times, int n, int k) {
        unordered_map<int, vector<Edge>> graph;
        for (const auto& edge : times) {
            int src = edge[0];
            int dst = edge[1];
            int time = edge[2];
            if (graph.find(src) == graph.end()) {
                graph[src] = vector<Edge>();
            }
            graph[src].push_back(Edge{dst, time});
        }

        vector<int> dist(n+1, INT_MAX);
        dist[k] = 0;

        int globalTime = 0;
        int visitedNum = 0;
        vector<bool> visitedArr(n+1, false);
        
        auto cmp = [](const Edge& a, const Edge& b) {
            return a.time > b.time;
        };
        priority_queue<Edge, vector<Edge>, decltype(cmp)> heap(cmp);
        heap.push(Edge{k, 0});

        for(;!heap.empty(); heap.pop()) {
            Edge cur = heap.top();
            if (visitedArr[cur.dst]) {
                continue;
            }

            visitedArr[cur.dst] = true;
            visitedNum++;
            globalTime = cur.time;

            for (const auto& edge : graph[cur.dst]) {
                int timeToDst = globalTime + edge.time;
                if (timeToDst < dist[edge.dst]) {
                    dist[edge.dst] = timeToDst;
                    heap.push(Edge{edge.dst, timeToDst});
                }
            }
        }
        if (visitedNum == n) {
            return globalTime;
        }
        return -1;
    }
};

struct test {
    vector<vector<int>> times;
    int n, k, expected;
};

int main() {
    vector<test> tests {
        {{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2, 2},
		{{{1, 2, 1}}, 2, 1, 1},
		{{{1, 2, 1}}, 2, 2, -1},
		{{{1, 2, 1}, {2, 3, 2}, {1, 3, 4}}, 3, 1, 3},
		{{{1, 2, 1}, {2, 3, 2}, {1, 3, 2}}, 3, 1, 2},
    };

    bool failed = false;
    int testOnly = 0;
    int i = 0;
    Solution sol;

    for (auto test : tests) {
        i++;
        if (testOnly && testOnly != i) {
            continue;
        }
        int actual = sol.networkDelayTime(test.times, test.n, test.k);
        if (actual != test.expected) {
            failed = true;
            cout << "Test " << i << " Failed\n\tActual  : " << actual << "\n\tExpected: " << test.expected << endl;
        }
    }
    if (!failed) {
        cout << "Passed All Tests." << endl;
    }
    return 0;
}