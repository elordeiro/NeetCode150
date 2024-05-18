#include <iostream>
#include <vector>
#include <string>
#include <unordered_set>
#include <algorithm> // Add this line

using namespace std;

class Solution {
public:
    bool wordBreak(string s, vector<string>& wordDict) {
        int n1 = s.length() + 1;
        unordered_set<string> wordSet(wordDict.begin(), wordDict.end());
        
        int shortest = (*min_element(wordDict.begin(), wordDict.end(), [](const std::string& a, const std::string& b) {
            return a.length() < b.length(); 
        })).length();
        int longest = (*max_element(wordDict.begin(), wordDict.end(), [](const std::string& a, const std::string& b) {
            return a.length() < b.length(); 
        })).length();

        vector<bool> dp(n1, false);
        dp[0] = true;
        for (int i = shortest; i < n1; i++) {
            for (int j = i - shortest; j > -1 && j > i - longest - 1; j--) {
                if (dp[j] && wordSet.find(s.substr(j, i - j)) != wordSet.end()) {
                    dp[i] = true;
                    break;
                }
            }
        }
        return dp[n1 - 1];
    }
};

struct test {
    string s;
    vector<string> wordDict;
    bool expected;
};

int main() {
    Solution sol;
    vector<test> tests {
        {"leetcode", {"leet","code"}, true},
        {"applepenapple", {"apple","pen"}, true},
        {"catsandog", {"cats","dog","sand","and","cat"}, false},
        {"bb", {"a","b","bbb","bbbb"}, true},
        {"aaaaaaa", {"aaaa","aaa"}, true},
        {"abcd", {"a","abc","b","cd"}, true},
    };

    bool failed = false;
    int testOnly = 0;
    int i = 0;
    for (auto test : tests) {
        if (testOnly && testOnly != i + 1) {
            continue;
        }
        i++;
        bool actual = sol.wordBreak(test.s, test.wordDict);
        if (actual != test.expected) {
            failed = true;
            cout << "Test " << i + 1 << " Failed\n\t" << "Actual  : " << actual << "\n\tExpected: " << test.expected << "\n";
        }
    }
    if (!failed) {
        cout << "Passed All Tests";
    }
    
    return 0;
}