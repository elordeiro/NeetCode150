package main

import (
	"fmt"
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}
	for _, str := range strs {
		slice := []byte(str)
		slices.Sort(slice)
		m[string(slice)] = append(m[string(slice)], str)
	}
	res := make([][]string, 0, len(m))
	for _, val := range m {
		res = append(res, val)
	}
	return res
}

func print_failed_test(i int, res [][]string, expected [][]string) {
	fmt.Printf("Test %d failed:\n", i+1)
	fmt.Printf("    \tActual  : %v\n", res)
	fmt.Printf("    \tExpected: %v\n", expected)
}

func main() {
	tests := []struct {
		strings  []string
		expected [][]string
	}{
		{
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			[]string{""},
			[][]string{{""}},
		},
		{
			[]string{"a"},
			[][]string{{"a"}},
		},
	}

	passed_all := true
	for i, test := range tests {
		strings, expected := test.strings, test.expected
		res := groupAnagrams(strings)
		if len(res) != len(expected) {
			passed_all = false
			print_failed_test(i, res, expected)
			continue
		}
		for j, str := range res {
			if len(str) != len(expected[j]) {
				passed_all = false
				print_failed_test(i, res, expected)
				break
			}
			for k, s := range str {
				if s != expected[j][k] {
					passed_all = false
					print_failed_test(i, res, expected)
					break
				}
			}
		}
	}
	if passed_all {
		fmt.Println("Passed all tests.")
	}
}
