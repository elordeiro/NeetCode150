package main

import (
	"fmt"
	"strconv"
)

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

func evalRPN(tokens []string) int {
	stack := newStack()

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			y := stack.pop()
			x := stack.pop()
			switch token {
			case "+":
				stack.push(x + y)
			case "-":
				stack.push(x - y)
			case "*":
				stack.push(x * y)
			case "/":
				stack.push(x / y)
			}
		default:
			num, _ := strconv.Atoi(token)
			stack.push(num)
		}
	}
	return stack.pop()
}

func print_failed_test(i int, res int, expected int) {
	fmt.Printf("Test %d failed:\n", i+1)
	fmt.Printf("    \tActual  : %v\n", res)
	fmt.Printf("    \tExpected: %v\n", expected)
}

func main() {
	tests := []struct {
		nums     []string
		expected int
	}{
		{[]string{"2", "1", "+", "3", "*"}, 9},
		{[]string{"4", "13", "5", "/", "+"}, 6},
		{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
		{
			[]string{
				"-78", "-33", "196", "+", "-19", "-", "115", "+", "-", "-99", "/", "-18", "8", "*", "-86", "-", "-", "16", "/", "26", "-14", "-", "-", "47", "-", "101", "-", "163", "*", "143", "-", "0", "-", "171", "+", "120", "*", "-60", "+", "156", "/", "173", "/", "-24", "11", "+", "21", "/", "*", "44", "*", "180", "70", "-40", "-", "*", "86", "132", "-84", "+", "*", "-", "38", "/", "/", "21", "28", "/", "+", "83", "/", "-31", "156", "-", "+", "28", "/", "95", "-", "120", "+", "8", "*", "90", "-", "-94", "*", "-73", "/", "-62", "/", "93", "*", "196", "-", "-59", "+", "187", "-", "143", "/", "-79", "-89", "+", "-",
			}, 165,
		},
	}

	passed_all := true
	for i, test := range tests {
		nums, expected := test.nums, test.expected
		res := evalRPN(nums)

		if res != expected {
			passed_all = false
			print_failed_test(i, res, expected)

		}
	}

	if passed_all {
		fmt.Println("Passed all tests.")
	}
}
