package main

import "fmt"

func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	n1 := len(s) + 1
	one_back, two_back := 1, 1

	for i := 2; i < n1; i++ {
		current := s[i-1]
		prev := s[i-2]

		temp := 0
		if current != '0' {
			temp += one_back
		}
		if prev == '1' || (prev == '2' && current < '7') {
			temp += two_back
		}
		two_back = one_back
		one_back = temp
	}
	return one_back
}

func main() {
	tests := []struct {
		s        string
		expected int
	}{
		{"12", 2},
		{"226", 3},
		{"06", 0},
		{"10", 1},
		{"27", 1},
		{"2101", 1},
	}
	failed := false
	testOnly := 0
	for i, test := range tests {
		if testOnly != 0 && testOnly != i+1 {
			continue
		}
		actual := numDecodings(test.s)
		if actual != test.expected {
			failed = true
			fmt.Printf("Test %v Failed\nActual  : %v\n\tExpected: %v\n", i+1, actual, test.expected)
		}
	}
	if !failed {
		fmt.Println("Passed All Tests.")
	}
}
