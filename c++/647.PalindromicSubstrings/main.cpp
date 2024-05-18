#include <string>
#include <vector>
#include <iostream>

using namespace std;

class Solution {
public:
    int countSubstrings(string s) {
        int count = 0;
        int n = s.length();
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < 2; j++) {
                int l = i;
                int r = i + j;
                while (l > -1 && r < n && s[l] == s[r]) {
                    count++;
                    l--;
                    r++;
                }
            }
        }
        return count;
    }
};

int main () {

    Solution sol;
    vector<pair<string, int> > tests = {
        {"abc", 3},
        {"aaa", 6},
        {"aba", 4},
        {"dnncbwoneinoplypwgbwktmvkoimcooyiwirgbxlcttgteqthcvyoueyftiwgwwxvxvg", 77}
    };

    bool failed = false;
    int testOnly = 0;
    int i = 0;
    for (auto test : tests) {
        if (testOnly != 0 && testOnly != i + 1) {
            continue;
        }
        string s = test.first;
        int expected = test.second;
        int actual = sol.countSubstrings(s);
        if (actual != expected) {
            failed = true;
            cout << "Test " << i + 1 << " Failed\n\tActual  :" << actual << "\n\tExpected: " << expected << "\n";
        }
    }
    if (!failed) {
        cout << "All Tests Passed.\n";
    }

}