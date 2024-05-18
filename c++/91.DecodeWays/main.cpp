#include <iostream>
#include <string>
#include <vector>

using namespace std;

class Solution {
public:
    int numDecodings(string s) {
        if (s.empty() || s[0] == '0') {
            return 0;
        }
        int n1 = s.size() + 1;
        int one_back = 1;
        int two_back = 1;
    
        for (int i = 2; i < n1; i++) {
            int current = s[i - 1];
            int prev = s[i - 2];

            int temp = 0;
            if (current != '0') {
                temp += one_back;
            }
            if (prev == '1' || (prev == '2' && current <= '6')) {
                temp += two_back;
            }
            
            two_back = one_back;
            one_back = temp;
        }
        return one_back;
    }
};

int main() {
    Solution s;
    vector<pair<string, int> > tests = {
        {"12", 2},
        {"226", 3},
        {"06", 0},
        {"10", 1},
        {"27", 1},
        {"2101", 1},
    };

    bool failed = false;
    int testOnly = 0;
    for (int i = 0; i < tests.size(); i++) {
        if (testOnly != 0 && testOnly != i + 1) {
            continue;
        }
        auto test = tests[i];
        auto result = s.numDecodings(test.first);
        if (result != test.second) {
            failed = true;
            cout << "Test: " << test.first << ", expect: " << test.second << ", while returned: " << result << endl;
        }
    }
    if (!failed) {
        cout << "All tests passed!" << endl;
    }

    return 0;
}