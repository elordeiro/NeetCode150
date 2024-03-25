package main

import "fmt"

type Stack struct {
	stack []int
	len   int
	cap   int
}

func newStack() *Stack {
	return &Stack{
		stack: make([]int, 4, 4),
		len:   0,
		cap:   4,
	}
}

func (stk *Stack) push(val int) {
	if stk.len == stk.cap {
		stk.stack = append(stk.stack, make([]int, stk.len, stk.cap)...)
		stk.cap *= 2
	}
	stk.stack[stk.len] = val
	stk.len++
}

func (stk *Stack) pop() int {
	res := stk.stack[stk.len-1]
	stk.len--
	if stk.len < stk.cap/4 {
		stk.stack = stk.stack[:stk.cap/2]
		stk.cap /= 2
	}
	return res
}

func (stk *Stack) isEmpty() bool {
	return stk.len == 0
}

func (stk *Stack) peek() int {
	return stk.stack[stk.len-1]
}

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures), len(temperatures))
	stack := newStack()
	for d, t := range temperatures {
		for !stack.isEmpty() && temperatures[stack.peek()] < t {
			pd := stack.pop()
			res[pd] = d - pd
		}
		stack.push(d)
	}
	return res
}

func print_failed_test(i int, res []int, expected []int) {
	fmt.Printf("Test %d failed:\n", i+1)
	fmt.Printf("    \tActual  : %v\n", res)
	fmt.Printf("    \tExpected: %v\n", expected)
}

func main() {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{[]int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
		{[]int{30, 60, 90}, []int{1, 1, 0}},
	}

	passed_all := true
	for i, test := range tests {
		nums, expected := test.nums, test.expected
		res := dailyTemperatures(nums)
		if len(res) != len(expected) {
			passed_all = false
			print_failed_test(i, res, expected)
			continue
		}
		for j, num := range res {
			if num != expected[j] {
				passed_all = false
				print_failed_test(i, res, expected)
				break
			}
		}
	}
	if passed_all {
		fmt.Println("Passed all tests.")
	}
}
