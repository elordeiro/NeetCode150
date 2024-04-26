package main

import (
	"fmt"
	"reflect"
	"testing"
)

const none = false

func Test211DesignAddAndSearchWordsDataStructure(t *testing.T) {
	tests := []struct {
		ops      []string
		args     [][]string
		expected []bool
	}{

		{
			[]string{"WordDictionary", "addWord", "addWord", "addWord", "search", "search", "search", "search"},
			[][]string{{}, {"bad"}, {"dad"}, {"mad"}, {"pad"}, {"bad"}, {".ad"}, {"b.."}},
			[]bool{none, none, none, none, false, true, true, true},
		},
		{
			[]string{"WordDictionary", "addWord", "addWord", "search", "search", "search", "search", "search", "search"},
			[][]string{{}, {"a"}, {"a"}, {"."}, {"a"}, {"aa"}, {"a"}, {".a"}, {"a."}},
			[]bool{none, none, none, true, true, false, true, false, false},
		},
	}
	test_only := 0
	for i, test := range tests {
		if test_only != 0 && test_only != i+1 {
			continue
		}
		testname := "Test211DesignAddAndSearchWordsDataStructure " + fmt.Sprint(i+1)
		actual := []bool{}
		var obj WordDictionary
		for j, op := range test.ops {
			if op == "WordDictionary" {
				obj = Constructor()
				actual = append(actual, none)
			} else if op == "addWord" {
				obj.AddWord(test.args[j][0])
				actual = append(actual, none)
			} else if op == "search" {
				actual = append(actual, obj.Search(test.args[j][0]))
			}
		}
		t.Run(testname, func(t *testing.T) {
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("\n\tActual  : %v\n\tExpected: %v", actual, test.expected)
			}
		})
	}
}
