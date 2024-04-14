package utils

type stack struct {
	stack []any
	len   int
	cap   int
}

func Stack() *stack {
	return &stack{
		stack: make([]any, 4, 4),
		len:   0,
		cap:   4,
	}
}

func (stk *stack) Push(val any) {
	if stk.len == stk.cap {
		stk.stack = append(stk.stack, make([]any, stk.len, stk.cap)...)
		stk.cap *= 2
	}
	stk.stack[stk.len] = val
	stk.len++
}

func (stk *stack) Pop() any {
	res := stk.stack[stk.len-1]
	stk.len--
	if stk.len < stk.cap/4 {
		stk.stack = stk.stack[:stk.cap/2]
		stk.cap /= 2
	}
	return res
}

func (stk *stack) IsEmpty() bool {
	return stk.len == 0
}

func (stk *stack) Peek() any {
	return stk.stack[stk.len-1]
}
