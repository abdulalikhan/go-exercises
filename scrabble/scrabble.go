// Package scrabble is a solution to the "Scrabble Score" problem
// https://exercism.org/tracks/go/exercises/scrabble-score/
package scrabble

import "strings"

// Score computes the Scrabble score for a given word
func Score(word string) int {
	// scoreMap maps each letter to a score
	var scoreMap = map[byte]int{
		'A': 1, 'B': 3, 'C': 3, 'D': 2, 'E': 1,
		'F': 4, 'G': 2, 'H': 4, 'I': 1, 'J': 8,
		'K': 5, 'L': 1, 'M': 3, 'N': 1, 'O': 1,
		'P': 3, 'Q': 10, 'R': 1, 'S': 1, 'T': 1,
		'U': 1, 'V': 4, 'W': 4, 'X': 8, 'Y': 4,
		'Z': 10,
	}
	// transform the word to uppercase
	word = strings.ToUpper(word)
	var score int = 0
	for i := 0; i < len(word); i++ {
		score += scoreMap[word[i]]
	}
	return score
}
