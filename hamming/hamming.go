// Package hamming is a solution to the "Hamming" problem
// https://exercism.org/tracks/go/exercises/hamming/
package hamming

import "errors"

// Distance calculates the Hamming Distance between two DNA strands represented as strings
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("the length of the sequences is not equal")
	}
	var count int = 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}
	return count, nil
}
