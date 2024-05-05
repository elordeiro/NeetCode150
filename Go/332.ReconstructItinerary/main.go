package main

import (
	"fmt"
	"reflect"
	"slices"
	"sort"
)

func findItinerary(tickets [][]string) []string {
	sort.Slice(tickets, func(i, j int) bool {
		for x := range tickets[i] {
			if tickets[i][x] == tickets[j][x] {
				continue
			}
			return tickets[i][x] < tickets[j][x]
		}
		return false
	})

	flights := len(tickets)
	fromSrcToDst := make(map[string][]string)
	for _, ticket := range tickets {
		src, dst := ticket[0], ticket[1]
		if _, ok := fromSrcToDst[src]; !ok {
			fromSrcToDst[src] = make([]string, 0)
		}
		fromSrcToDst[src] = append(fromSrcToDst[src], dst)
	}

	res := []string{"JFK"}
	var dfs func(string) bool
	dfs = func(src string) bool {
		if flights == 0 {
			return true
		}
		failed := make(map[string]struct{})
		for i, dst := range fromSrcToDst[src] {
			if _, ok := failed[dst]; ok {
				continue
			}
			fromSrcToDst[src] = slices.Concat(fromSrcToDst[src][:i], fromSrcToDst[src][i+1:])
			res = append(res, dst)
			flights -= 1

			if dfs(dst) {
				return true
			}

			failed[dst] = struct{}{}
			fromSrcToDst[src] = slices.Concat(fromSrcToDst[src][:i], []string{dst}, fromSrcToDst[src][i:])
			flights += 1
			res = res[:len(res)-1]
		}
		return false
	}
	dfs("JFK")
	return res
}

func main() {
	tests := []struct {
		tickets  [][]string
		expected []string
	}{
		{[][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}, []string{"JFK", "MUC", "LHR", "SFO", "SJC"}},
		{[][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}, []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"}},
		{[][]string{{"EZE", "AXA"}, {"TIA", "ANU"}, {"ANU", "JFK"}, {"JFK", "ANU"}, {"ANU", "EZE"}, {"TIA", "ANU"}, {"AXA", "TIA"}, {"TIA", "JFK"}, {"ANU", "TIA"}, {"JFK", "TIA"}}, []string{"JFK", "ANU", "EZE", "AXA", "TIA", "ANU", "JFK", "TIA", "ANU", "TIA", "JFK"}},
		{[][]string{{"AXA", "EZE"}, {"EZE", "AUA"}, {"ADL", "JFK"}, {"ADL", "TIA"}, {"AUA", "AXA"}, {"EZE", "TIA"}, {"EZE", "TIA"}, {"AXA", "EZE"}, {"EZE", "ADL"}, {"ANU", "EZE"}, {"TIA", "EZE"}, {"JFK", "ADL"}, {"AUA", "JFK"}, {"JFK", "EZE"}, {"EZE", "ANU"}, {"ADL", "AUA"}, {"ANU", "AXA"}, {"AXA", "ADL"}, {"AUA", "JFK"}, {"EZE", "ADL"}, {"ANU", "TIA"}, {"AUA", "JFK"}, {"TIA", "JFK"}, {"EZE", "AUA"}, {"AXA", "EZE"}, {"AUA", "ANU"}, {"ADL", "AXA"}, {"EZE", "ADL"}, {"AUA", "ANU"}, {"AXA", "EZE"}, {"TIA", "AUA"}, {"AXA", "EZE"}, {"AUA", "SYD"}, {"ADL", "JFK"}, {"EZE", "AUA"}, {"ADL", "ANU"}, {"AUA", "TIA"}, {"ADL", "EZE"}, {"TIA", "JFK"}, {"AXA", "ANU"}, {"JFK", "AXA"}, {"JFK", "ADL"}, {"ADL", "EZE"}, {"AXA", "TIA"}, {"JFK", "AUA"}, {"ADL", "EZE"}, {"JFK", "ADL"}, {"ADL", "AXA"}, {"TIA", "AUA"}, {"AXA", "JFK"}, {"ADL", "AUA"}, {"TIA", "JFK"}, {"JFK", "ADL"}, {"JFK", "ADL"}, {"ANU", "AXA"}, {"TIA", "AXA"}, {"EZE", "JFK"}, {"EZE", "AXA"}, {"ADL", "TIA"}, {"JFK", "AUA"}, {"TIA", "EZE"}, {"EZE", "ADL"}, {"JFK", "ANU"}, {"TIA", "AUA"}, {"EZE", "ADL"}, {"ADL", "JFK"}, {"ANU", "AXA"}, {"AUA", "AXA"}, {"ANU", "EZE"}, {"ADL", "AXA"}, {"ANU", "AXA"}, {"TIA", "ADL"}, {"JFK", "ADL"}, {"JFK", "TIA"}, {"AUA", "ADL"}, {"AUA", "TIA"}, {"TIA", "JFK"}, {"EZE", "JFK"}, {"AUA", "ADL"}, {"ADL", "AUA"}, {"EZE", "ANU"}, {"ADL", "ANU"}, {"AUA", "AXA"}, {"AXA", "TIA"}, {"AXA", "TIA"}, {"ADL", "AXA"}, {"EZE", "AXA"}, {"AXA", "JFK"}, {"JFK", "AUA"}, {"ANU", "ADL"}, {"AXA", "TIA"}, {"ANU", "AUA"}, {"JFK", "EZE"}, {"AXA", "ADL"}, {"TIA", "EZE"}, {"JFK", "AXA"}, {"AXA", "ADL"}, {"EZE", "AUA"}, {"AXA", "ANU"}, {"ADL", "EZE"}, {"AUA", "EZE"}}, []string{"JFK", "ADL", "ANU", "ADL", "ANU", "AUA", "ADL", "AUA", "ADL", "AUA", "ANU", "AXA", "ADL", "AUA", "ANU", "AXA", "ADL", "AXA", "ADL", "AXA", "ANU", "AXA", "ANU", "AXA", "EZE", "ADL", "AXA", "EZE", "ADL", "AXA", "EZE", "ADL", "EZE", "ADL", "EZE", "ADL", "EZE", "ANU", "EZE", "ANU", "EZE", "AUA", "AXA", "EZE", "AUA", "AXA", "EZE", "AUA", "AXA", "JFK", "ADL", "EZE", "AUA", "EZE", "AXA", "JFK", "ADL", "JFK", "ADL", "JFK", "ADL", "JFK", "ADL", "TIA", "ADL", "TIA", "AUA", "JFK", "ANU", "TIA", "AUA", "JFK", "AUA", "JFK", "AUA", "TIA", "AUA", "TIA", "AXA", "TIA", "EZE", "AXA", "TIA", "EZE", "JFK", "AXA", "TIA", "EZE", "JFK", "AXA", "TIA", "JFK", "EZE", "TIA", "JFK", "EZE", "TIA", "JFK", "TIA", "JFK", "AUA", "SYD"}},
		{[][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "JFK"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}}, []string{"JFK", "SFO", "JFK", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL"}},
	}

	testOnly := 0
	for i, test := range tests {
		if testOnly != 0 && testOnly != i+1 {
			continue
		}
		actual := findItinerary(test.tickets)
		if !reflect.DeepEqual(actual, test.expected) {
			fmt.Printf("Test %v Failed\n\tActual  : %v\n\tExpected: %v\n", i+1, actual, test.expected)
		}
	}

	// i := 6
	// arr := []int{0, 1, 2, 3, 4, 5, 6}
	// num := arr[i]
	// arr = slices.Concat(arr[:i], arr[i+1:])
	// fmt.Printf("Num: %v\nArr: %v\n", num, arr)
	// arr = slices.Concat(arr[:i], []int{num}, arr[i:])
	// fmt.Printf("Num: %v\nArr: %v\n", num, arr)
}
