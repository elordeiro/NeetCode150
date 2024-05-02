package main

import (
	"fmt"
	"reflect"
)

func findOrder(numCourses int, prerequisites [][]int) []int {
    prereqs := make([][]int, numCourses)
	for _, edge := range prerequisites {
		course, prereq := edge[0], edge[1]
		prereqs[course] = append(prereqs[course], prereq)
	}
	
	need := make(map[int]struct{})
	taken := make(map[int]struct{})
	res := []int{}

	var hasCycle func(int) bool
	hasCycle = func(course int) bool {
		if _, ok := need[course]; ok {
			return true
		}
		if _, ok := taken[course]; ok {
			return false
		}
		need[course] = struct{}{}
		for _, c := range prereqs[course] {
			if hasCycle(c) {
				return true
			}
		}
		delete(need, course)
		taken[course] = struct{}{}
		res = append(res, course)
		return false
	}

	for course := range numCourses {
		if hasCycle(course) {
			return []int{}
		}
	}
	return res
}

func main() {
	tests := []struct{
		numCourses int
		prerequisites [][]int
		expected []int
	}{
		{2, [][]int{{1,0}}, []int{0,1}},
        {2, [][]int{{0,1}}, []int{1,0}},
        {4, [][]int{{1,0},{2,0},{3,1},{3,2}}, []int{0,2,1,3}},
        {1, [][]int{}, []int{0}},
	}

	testOnly := 0
	for i, test := range tests {
		if testOnly != 0 && testOnly != i + 1{
			continue
		}
		actual := findOrder(test.numCourses, test.prerequisites)
		if !reflect.DeepEqual(actual, test.expected) {
			fmt.Printf("Test %v Failed\n\tActual  : %v\n\tExpected: %v\n\t", i + 1, actual, test.expected)
		}
	}
}

