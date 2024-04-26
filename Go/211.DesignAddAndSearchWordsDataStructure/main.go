package main

type WordDictionary struct {
	children [27]*WordDictionary
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	curr := this
	for _, char := range word {
		char -= 'a'
		if curr.children[char] == nil {
			curr.children[char] = &WordDictionary{}
		}
		curr = curr.children[char]

	}
	curr.children[26] = &WordDictionary{}
}

func (this *WordDictionary) dfs(i int, word string) bool {
	if this == nil {
		return false
	}
	if i == len(word) {
		return this.children[26] != nil
	}
	char := word[i]
	if char == '.' {
		for _, child := range this.children {
			if child.dfs(i+1, word) {
				return true
			}
		}
	} else if char -= 'a'; this.children[char] != nil {
		return this.children[char].dfs(i+1, word)
	}
	return false
}

func (this *WordDictionary) Search(word string) bool {
	return this.dfs(0, word)
}
