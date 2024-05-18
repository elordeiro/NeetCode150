#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    int coinChange(vector<int>& coins, int amount) {
        vector<int> memo(amount + 1, 10001);
        memo[0] = 0;
        for (int i = 1; i < amount + 1; i++) {
            int min_change = 10001;
            for (auto coin : coins) {
                if (i == coin) {
                    min_change = 1;
                    break;
                }
                if (i - coin > -1) {
                    min_change = min(min_change, memo[i-coin] + 1);
                }
            }
            memo[i] = min_change;
        }
        if (memo[amount] != 10001) {
            return memo[amount];
        }
        return -1;
    }
};


int main() {
    Solution sol;
    vector<tuple<vector<int>, int, int> > tests = {
        {{1,2,5}, 11, 3},
        {{2}, 3, -1},
        {{1}, 0, 0},
        {{2,5,10,1}, 27, 4},
        {{186,419,83,408}, 6249, 20},
    };
    bool failed = false;
    int testOnly = 0;
    for (int i = 0; i < tests.size(); i++) {
        if (testOnly != 0 && testOnly != i + 1) {
            continue;
        }
        auto test = tests[i];
        auto coins = get<0>(test);
        auto amount = get<1>(test);
        auto expected = get<2>(test);
        auto actual = sol.coinChange(coins, amount);
        if (actual != expected) {
            failed = true;
            cout << "Test " << i + 1 << " Failed\n\tActual  :" << actual << "\n\tExpected: " << expected << "\n";
        }
    }
    if (!failed) {
        cout << "All Tests Passed.\n";
    }
}