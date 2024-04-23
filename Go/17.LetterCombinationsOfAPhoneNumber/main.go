package main

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	res := []string{}
	n := len(digits)
	keypad := map[byte]string{'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}
	var backtrack func(int, string)
	backtrack = func(i int, partial string) {
		if i == n {
			res = append(res, partial)
			return
		}

		for _, char := range keypad[digits[i]] {
			partial += string(char)
			backtrack(i+1, partial)
			partial = partial[:len(partial)-1]
		}
	}
	backtrack(0, "")
	return res
}
