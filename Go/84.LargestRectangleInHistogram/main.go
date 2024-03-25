package main

import "fmt"

type Stack struct {
	stack [][2]int
	len   int
	cap   int
}

func newStack() *Stack {
	return &Stack{
		stack: make([][2]int, 4, 4),
		len:   0,
		cap:   4,
	}
}

func (stk *Stack) push(val [2]int) {
	if stk.len == stk.cap {
		stk.stack = append(stk.stack, make([][2]int, stk.len, stk.cap)...)
		stk.cap *= 2
	}
	stk.stack[stk.len] = val
	stk.len++
}

func (stk *Stack) pop() [2]int {
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

func (stk *Stack) peek() [2]int {
	return stk.stack[stk.len-1]
}

func largestRectangleArea(heights []int) int {
	stack := newStack()
	largest := 0
	for j, h := range heights {
		i := j
		for !stack.isEmpty() && stack.peek()[1] > h {
			i = stack.peek()[0]
			height := stack.pop()[1]
			largest = max(largest, height*(j-i))
		}
		stack.push([2]int{i, h})
	}
	j := len(heights)
	for !stack.isEmpty() {
		i, height := stack.peek()[0], stack.pop()[1]
		largest = max(largest, height*(j-i))
	}
	return largest
}

func print_failed_test(i int, res int, expected int) {
	fmt.Printf("Test %d failed:\n", i+1)
	fmt.Printf("    \tActual  : %v\n", res)
	fmt.Printf("    \tExpected: %v\n", expected)
}

func main() {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{2, 1, 5, 6, 2, 3}, 10},
		{[]int{2, 4}, 4},
		{[]int{2, 1, 2}, 3},
	}

	passed_all := true
	for i, test := range tests {
		nums, expected := test.nums, test.expected
		res := largestRectangleArea(nums)
		if res != expected {
			passed_all = false
			print_failed_test(i, res, expected)
		}
	}
	if passed_all {
		fmt.Println("Passed all tests.")
	}
}
