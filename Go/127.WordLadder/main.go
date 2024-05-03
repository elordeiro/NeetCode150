package main

import (
	"container/list"
	"fmt"
)

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]struct{})
	for _, word := range wordList {
		wordSet[word] = struct{}{}
	}

	charMap := make(map[int]map[rune]struct{})
	for _, word := range wordList {
		for i, char := range word {
			if _, ok := charMap[i]; !ok {
				charMap[i] = make(map[rune]struct{})
			}
			charMap[i][char] = struct{}{}
		}
	}

	steps := 1
	fVisited := make(map[string]struct{})
	bVisited := make(map[string]struct{})
	fVisited[beginWord] = struct{}{}
	bVisited[endWord] = struct{}{}
	fQueue := list.New()
	bQueue := list.New()
	fQueue.PushBack(beginWord)
	bQueue.PushBack(endWord)
	for fQueue.Len() > 0 && bQueue.Len() > 0 {
		if bQueue.Len() > fQueue.Len() {
			fQueue, bQueue = bQueue, fQueue
			fVisited, bVisited = bVisited, fVisited
		}
		for range fQueue.Len() {
			word := fQueue.Remove(fQueue.Front()).(string)
			wordRunes := []rune(word)
			for i := range wordRunes {
				old := wordRunes[i]
				for char := range charMap[i] {
					wordRunes[i] = char
					adj := string(wordRunes)
					if _, inWordSet := wordSet[adj]; inWordSet {
						if _, inFvisited := fVisited[adj]; !inFvisited {
							if _, found := bVisited[adj]; found {
								return steps + 1
							}
							fQueue.PushBack(adj)
							fVisited[adj] = struct{}{}
						}
					}
				}
				wordRunes[i] = old
			}
		}
		steps++
	}
	return 0
}

func main() {
	tests := []struct {
		beginWord string
		endWord   string
		wordList  []string
		expected  int
	}{
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}, 5},
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}, 0},
	}
	for i, test := range tests {
		actual := ladderLength(test.beginWord, test.endWord, test.wordList)
		if actual != test.expected {
			fmt.Printf("Test %v Failed\n\tActual  : %v\n\tExpected: %v\n", i+1, actual, test.expected)
		}
	}
}
