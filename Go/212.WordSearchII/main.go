package main

type TrieNode struct {
	children map[rune]*TrieNode
	word     string
}

func findWords(board [][]byte, words []string) []string {
	root := &TrieNode{children: map[rune]*TrieNode{}}
	for _, word := range words {
		curr := root
		for _, char := range word {
			if curr.children[char] == nil {
				curr.children[char] = &TrieNode{children: map[rune]*TrieNode{}}
			}
			curr = curr.children[char]
		}
		curr.word = word
	}

	res := []string{}
	m, n := len(board)-1, len(board[0])-1

	var backtrack func(int, int, *TrieNode)
	backtrack = func(row int, col int, root *TrieNode) {
		char := rune(board[row][col])
		_, ok := root.children[char]
		if char == '#' || !ok {
			return
		}
		child := root.children[char]
		if child.word != "" {
			res = append(res, child.word)
			child.word = ""
		}
		if len(child.children) == 0 {
			delete(root.children, char)
			return
		}
		board[row][col] = '#'
		if row < m {
			backtrack(row+1, col, child)
		}
		if col < n {
			backtrack(row, col+1, child)
		}
		if row > 0 {
			backtrack(row-1, col, child)
		}
		if col > 0 {
			backtrack(row, col-1, child)
		}
		board[row][col] = byte(char)
	}

	for row := range m + 1 {
		for col := range n + 1 {
			if _, ok := root.children[rune(board[row][col])]; ok {
				backtrack(row, col, root)

			}
		}
	}
	return res
}
