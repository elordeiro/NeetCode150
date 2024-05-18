#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    int maxProduct(vector<int>& nums) {
        int res = nums[0];
        int curMin = 1;
        int curMax = 1;
        int temp;
        for (const auto &num : nums) {
            temp = curMax;
            curMax = max(num, max(num * curMax, num * curMin)); 
            curMin = min(num, min(num * curMin, num * temp));
            res = max(res, max(curMin, curMax));
        }
        return res;
    }
};

struct test {
    vector<int> nums;
    int expected;
};

int main() {
    Solution sol;
    vector<test> tests {
        {{2,3,-2,4}, 6},
        {{-2,0,-1}, 0},
        {{-2,3,-4}, 24},
        {{0, 2}, 2},
        {{2,-5,-2,-4,3}, 24},
        {{-2}, -2},
    };
    bool failed = false;
    int testOnly = 0;
    int i = 0;
    for (auto test : tests) {
        if (testOnly != 0 && testOnly != i + 1) {
            continue;
        }
        i++;
        int actual = sol.maxProduct(test.nums);
        if (actual != test.expected) {
            failed = true;
            cout << "Test " << i + 1 << " Failed\n\t" << "Actual  : " << actual << "\n\tExpected: " << test.expected << "\n";
        }
    }
    if (!failed) {
        cout << "Passed All Tests.\n";
    }
    
}