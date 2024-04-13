package utils

import (
	"math/rand/v2"
	"testing"
)

func TestStack(t *testing.T) {
	goStack := []int{}
	customStack := NewStack()
	for range 10 {
		num := rand.N(100)
		goStack = append(goStack, num)
		customStack.push(num)
	}
	var testname string
	for !customStack.isEmpty() {
		num1 := goStack[len(goStack)-1]
		num2 := customStack.pop()
		testname = "Test Stack Pop"
		t.Run(testname, func(t *testing.T) {
			if num1 != num2 {
				t.Errorf("Actual   : %v", num1)
				t.Errorf("Expected : %v", num2)
			}
		})
		goStack = goStack[:len(goStack)-1]
	}
}
