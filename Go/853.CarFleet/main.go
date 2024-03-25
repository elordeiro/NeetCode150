package main

import (
	"fmt"
	"sort"
)

type Stack struct {
	stack []float32
	len   int
	cap   int
}

func newStack() *Stack {
	return &Stack{
		stack: make([]float32, 4, 4),
		len:   0,
		cap:   4,
	}
}

func (stk *Stack) push(val float32) {
	if stk.len == stk.cap {
		stk.stack = append(stk.stack, make([]float32, stk.len, stk.cap)...)
		stk.cap *= 2
	}
	stk.stack[stk.len] = val
	stk.len++
}

func (stk *Stack) pop() float32 {
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

func (stk *Stack) peek() float32 {
	return stk.stack[stk.len-1]
}

func carFleet(target int, position []int, speed []int) int {
	posTospeed := map[int]int{}
	for i, pos := range position {
		posTospeed[pos] = speed[i]
	}
	sort.Slice(position, func(i, j int) bool {
		return position[j] < position[i]
	})
	stack := newStack()
	for _, pos := range position {
		time := float32((target - pos)) / float32(posTospeed[pos])
		if stack.isEmpty() || stack.peek() < time {
			stack.push(time)
		}
	}
	return stack.len
}

func print_failed_test(i int, res int, expected int) {
	fmt.Printf("Test %d failed:\n", i+1)
	fmt.Printf("    \tActual  : %v\n", res)
	fmt.Printf("    \tExpected: %v\n", expected)
}

func main() {
	tests := []struct {
		target   int
		position []int
		speed    []int
		expected int
	}{
		{12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}, 3},
		{10, []int{3}, []int{3}, 1},
		{100, []int{0, 2, 4}, []int{4, 2, 1}, 1},
		{10, []int{6, 8}, []int{3, 2}, 2},
	}

	passed_all := true
	test_only := 4
	for i, test := range tests {
		if test_only != 0 && test_only != i+1 {
			continue
		}
		target, position, speed, expected := test.target, test.position, test.speed, test.expected
		res := carFleet(target, position, speed)

		if res != expected {
			passed_all = false
			print_failed_test(i, res, expected)

		}
	}
	if passed_all {
		fmt.Println("Passed all tests.")
	}
}
