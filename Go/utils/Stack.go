package main

type Stack struct {
	stack []int
	len   int
	cap   int
}

func newStack() *Stack {
	return &Stack{
		stack: make([]int, 4, 4),
		len:   0,
		cap:   2,
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
