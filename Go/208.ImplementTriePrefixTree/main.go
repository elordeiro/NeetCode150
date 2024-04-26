package main

type Trie struct {
	children [27]*Trie
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	for _, char := range word {
		char -= 'a'
		if this.children[char] == nil {
			this.children[char] = &Trie{}
		}
		this = this.children[char]
	}
	this.children[26] = &Trie{}
}

func (this *Trie) Search(word string) bool {
	for _, char := range word {
		char -= 'a'
		if this.children[char] == nil {
			return false
		}
		this = this.children[char]
	}
	return this.children[26] != nil
}

func (this *Trie) StartsWith(prefix string) bool {
	for _, char := range prefix {
		char -= 'a'
		if this.children[char] == nil {
			return false
		}
		this = this.children[char]
	}
	return true
}
