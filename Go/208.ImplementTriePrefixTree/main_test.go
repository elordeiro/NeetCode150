package main

import (
	"fmt"
	"reflect"
	"testing"
)

var null = false

func Test208ImplementTriePrefixTree(t *testing.T) {
	tests := []struct {
		ops      []string
		args     [][]string
		expected []bool
	}{
		{
			[]string{"Trie", "insert", "search", "search", "startsWith", "insert", "search"},
			[][]string{{}, {"apple"}, {"apple"}, {"app"}, {"app"}, {"app"}, {"app"}},
			[]bool{null, null, true, false, true, null, true},
		},
		{
			[]string{"Trie", "insert", "search", "search", "startsWith", "insert", "search", "insert", "search"},
			[][]string{{}, {"apple"}, {"apple"}, {"app"}, {"app"}, {"app"}, {"app"}, {"app"}, {"apple"}},
			[]bool{null, null, true, false, true, null, true, null, true},
		},
		{
			[]string{"Trie", "insert", "insert", "insert", "insert", "insert", "insert", "search", "search", "search", "search", "search", "search", "search", "search", "search", "startsWith", "startsWith", "startsWith", "startsWith", "startsWith", "startsWith", "startsWith", "startsWith", "startsWith"},
			[][]string{{}, {"app"}, {"apple"}, {"beer"}, {"add"}, {"jam"}, {"rental"}, {"apps"}, {"app"}, {"ad"}, {"applepie"}, {"rest"}, {"jan"}, {"rent"}, {"beer"}, {"jam"}, {"apps"}, {"app"}, {"ad"}, {"applepie"}, {"rest"}, {"jan"}, {"rent"}, {"beer"}, {"jam"}},
			[]bool{null, null, null, null, null, null, null, false, true, false, false, false, false, false, true, true, false, true, true, false, false, false, true, true, true},
		},
	}
	for i, test := range tests {
		testname := "Test208ImplementTriePrefixTree" + fmt.Sprint(i+1)
		actual := []bool{}
		var obj Trie
		for j, op := range test.ops {
			if op == "Trie" {
				obj = Constructor()
				actual = append(actual, null)
			} else if op == "insert" {
				obj.Insert(test.args[j][0])
				actual = append(actual, null)
			} else if op == "startsWith" {
				actual = append(actual, obj.StartsWith(test.args[j][0]))
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
