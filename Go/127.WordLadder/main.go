package main

import (
	"fmt"
)

var wordSet map[string]struct{}
var charMap map[int]map[rune]struct{}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if beginWord == endWord {
		return 1
	}

	wordSet = make(map[string]struct{})
	charMap = make(map[int]map[rune]struct{})

	// Convert list into a set for constant time look up
	for _, word := range wordList {
		wordSet[word] = struct{}{}

		// Map indices to a 'set' of letters that can occur at that index
		for i, char := range word {
			if _, ok := charMap[i]; !ok {
				charMap[i] = make(map[rune]struct{})
			}
			charMap[i][char] = struct{}{}
		}
	}

	// Early return if set doesn't have the endWord
	if _, ok := wordSet[endWord]; !ok {
		return 0
	}

	// Set up Bidirectional BFS. We need a forward and backward queue
	// We have found the path if the queues ever have an overlapping value
	steps := 1
	fVisited := make(map[string]struct{})
	bVisited := make(map[string]struct{})
	fVisited[beginWord] = struct{}{}
	bVisited[endWord] = struct{}{}
	// I'll be using a slice rather than a queue for faster access so I also need some indices
	fQueue := []string{beginWord}
	bQueue := []string{endWord}
	fIdx, bIdx := 0, 0

	// Run Bidirectional BFS
	for fIdx < len(fQueue) && bIdx < len(bQueue) {
		var found bool

		// Forward BFS
		found, fQueue, fIdx = bfs(fQueue, fIdx, fVisited, bVisited)
		if found {
			return steps + 1
		}
		// One level done
		steps++

		// Backward BFS
		found, bQueue, bIdx = bfs(bQueue, bIdx, bVisited, fVisited)
		if found {
			return steps + 1
		}
		// Another level done
		steps++
	}
	return 0
}

func bfs(queue []string, idx int, visited, otherVisited map[string]struct{}) (bool, []string, int) {
	for range len(queue) - idx {
		// 'Popleft' from the queue
		word := queue[idx]
		idx++

		// Convert the string into a slice of runes so we can modify it
		wordRunes := []rune(word)

		// Modify each letter in the current word
		for i := range wordRunes {

			// Remember old value since we'll be creating the neighbors using this slice
			old := wordRunes[i]

			// Iterate over all chars in the set of this index
			for char := range charMap[i] {

				// Create a new word with every letter that can be at this position
				wordRunes[i] = char
				adj := string(wordRunes)

				// If the word is valid
				if _, inWordSet := wordSet[adj]; inWordSet {

					// And we have not visited it in forward direction
					if _, inVisited := visited[adj]; !inVisited {

						// If we have visited it in backward direction, path is found
						if _, found := otherVisited[adj]; found {
							return true, nil, 0
						}

						// Otherwise visit this neighbor
						queue = append(queue, adj)
						visited[adj] = struct{}{}
					}
				}
			}
			// Restore current word
			wordRunes[i] = old
		}
	}
	// Path no yet found
	return false, queue, idx
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
