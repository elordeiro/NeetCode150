package main

import (
	"fmt"
	"slices"
)

type TimeMap struct {
	keyToTime   map[string][]int
	timeToValue map[int]string
}

func Constructor() TimeMap {
	tm := new(TimeMap)
	tm.keyToTime = make(map[string][]int)
	tm.timeToValue = make(map[int]string)
	return *tm
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.keyToTime[key] = append(this.keyToTime[key], timestamp)
	this.timeToValue[timestamp] = value
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if _, ok := this.keyToTime[key]; !ok {
		return ""
	}
	tss := this.keyToTime[key]
	n := len(tss)
	l, r := 0, n-1
	for l <= r {
		m := (l + r) / 2
		if tss[m] < timestamp {
			l = m + 1
		} else if timestamp < tss[m] {
			r = m - 1
		} else {
			return this.timeToValue[timestamp]
		}
	}
	if r < 0 {
		return ""
	}
	i := min(l, n-1)
	for tss[i] > timestamp {
		i -= 1
	}
	return this.timeToValue[tss[i]]
}

func print_failed_test(i int, res []string, expected []string) {
	fmt.Printf("Test %d failed:\n", i+1)
	fmt.Printf("    \tActual  : %v\n", res)
	fmt.Printf("    \tExpected: %v\n", expected)
}

func main() {
	type test struct {
		actual   []string
		expected []string
	}
	tests := []test{}
	tm := Constructor()
	actual := []string{}
	tm.Set("foo", "bar", 1)
	actual = append(actual, tm.Get("foo", 1))
	actual = append(actual, tm.Get("foo", 3))
	tm.Set("foo", "bar2", 4)
	actual = append(actual, tm.Get("foo", 4))
	actual = append(actual, tm.Get("foo", 5))
	tests = append(tests, test{actual, []string{"bar", "bar", "bar2", "bar2"}})

	passed_all := true
	for i, test := range tests {
		actual, expected := test.actual, test.expected
		if !slices.Equal(actual, expected) {
			passed_all = false
			print_failed_test(i, actual, expected)
		}
	}
	if passed_all {
		fmt.Println("Passed all tests.")
	}
}
