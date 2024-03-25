package main

import (
	"fmt"
	"slices"
)

type MinStack struct {
	lastIdx int
	stack   [][2]int
}

func Constructor() MinStack {
	return MinStack{0, [][2]int{{int(2e18), int(2e18)}}}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, [2]int{val, min(val, this.stack[this.lastIdx][1])})
	this.lastIdx++
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:this.lastIdx]
	this.lastIdx--
}

func (this *MinStack) Top() int {
	return this.stack[this.lastIdx][0]
}

func (this *MinStack) GetMin() int {
	return this.stack[this.lastIdx][1]
}

func print_failed_test(i int, res []int, expected []int) {
	fmt.Printf("Test %d failed:\n", i+1)
	fmt.Printf("    \tActual  : %v\n", res)
	fmt.Printf("    \tExpected: %v\n", expected)
}

func main() {
	type test struct {
		actual   []int
		expected []int
	}
	tests := []test{}

	actual := []int{}
	ms := Constructor()
	ms.Push(-2)
	ms.Push(0)
	ms.Push(-3)
	actual = append(actual, ms.GetMin())
	ms.Pop()
	actual = append(actual, ms.Top())
	actual = append(actual, ms.GetMin())
	expected := []int{-3, 0, -2}
	tests = append(tests, test{actual, expected})

	actual = []int{}
	ms = Constructor()
	ms.Push(0)
	ms.Push(1)
	ms.Push(0)
	actual = append(actual, ms.GetMin())
	ms.Pop()
	actual = append(actual, ms.GetMin())
	expected = []int{0, 0}
	tests = append(tests, test{actual, expected})

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
