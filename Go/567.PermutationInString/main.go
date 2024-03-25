package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	var n, m int
	if n, m = len(s1), len(s2); n > m {
		return false
	}
	s1Count := make([]int, 26, 26)
	s2Count := make([]int, 26, 26)
	for i := range n {
		s1Count[s1[i]-97]++
		s2Count[s2[i]-97]++
	}

	matches := 0
	for i := range 26 {
		if s1Count[i] == s2Count[i] {
			matches++
		}
	}
	i, j := 0, n
	for j < m {
		if matches == 26 {
			return true
		}

		idx := s2[i] - 97
		s2Count[idx]--
		if s1Count[idx] == s2Count[idx] {
			matches++
		} else if s1Count[idx]-1 == s2Count[idx] {
			matches--
		}

		idx = s2[j] - 97
		s2Count[idx]++
		if s1Count[idx] == s2Count[idx] {
			matches++
		} else if s1Count[idx]+1 == s2Count[idx] {
			matches--
		}
		i++
		j++
	}
	return matches == 26
}

func print_failed_test(i int, res bool, expected bool) {
	fmt.Printf("Test %d failed:\n", i+1)
	fmt.Printf("    \tActual  : %v\n", res)
	fmt.Printf("    \tExpected: %v\n", expected)
}

func main() {
	tests := []struct {
		s1       string
		s2       string
		expected bool
	}{
		{"ab", "eidbaooo", true},
		{"ab", "eidboaoo", false},
		{"adc", "dcda", true},
		{"ab", "a", false},
		{"abc", "bbbca", true},
		{"trinitrophenylmethylnitramine", "dinitrophenylhydrazinetrinitrophenylmethylnitramine", true},
	}

	passed_all := true
	for i, test := range tests {
		s1, s2, expected := test.s1, test.s2, test.expected
		res := checkInclusion(s1, s2)

		if res != expected {
			passed_all = false
			print_failed_test(i, res, expected)
		}
	}
	if passed_all {
		fmt.Println("Passed all tests.")
	}
}
