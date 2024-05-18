package main

import "fmt"

func longestPalindrome(s string) string {
	n := len(s)
	res := ""

	expand := func(l, r int) {
		for l > -1 && r < n && s[l] == s[r] {
			if r-l+1 > len(res) {
				res = s[l : r+1]
			}
			l--
			r++
		}
	}

	for i := range n {
		expand(i, i)
		expand(i, i+1)
	}

	return res
}

func main() {
	tests := []struct {
		s, expected string
	}{
		{"babad", "bab"},
		{"cbbd", "bb"},
		{"a", "a"},
		{"babaddtattarrattatddetartrateedredividerb", "ddtattarrattatdd"},
		{"bbbb", "bbbb"},
		{"slvafhpfjpbqbpcuwxuexavyrtymfydcnvvbvdoitsvumbsvoayefsnusoqmlvatmfzgwlhxtkhdnlmqmyjztlytoxontggyytcezredlrrimcbkyzkrdeshpyyuolsasyyvxfjyjzqksyxtlenaujqcogpqmrbwqbiaweacvkcdxyecairvvhngzdaujypapbhctaoxnjmwhqdzsvpyixyrozyaldmcyizilrmmmvnjbyhlwvpqhnnbausoyoglvogmkrkzppvexiovlxtmustooahwviluumftwnzfbxxrvijjyfybvfnwpjjgdudnyjwoxavlyiarjydlkywmgjqeelrohrqjeflmdyzkqnbqnpaewjdfmdyoazlznzthiuorocncwjrocfpzvkcmxdopisxtatzcpquxyxrdptgxlhlrnwgvee", ""},
	}

	testOnly := 0
	failed := false
	for i, test := range tests {
		if testOnly != 0 && testOnly != i+1 {
			continue
		}
		actual := longestPalindrome(test.s)
		if actual != test.expected {
			failed = true
			fmt.Printf("Test %v Failed\n\tActual  : %v\n\tExpected: %v\n", i+1, actual, test.expected)
		}
	}
	if !failed {
		fmt.Println("Passed All Tests.")
	}

}
