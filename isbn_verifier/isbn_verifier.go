// Package isbn is a solution to the "ISBN Verifier" problem
// https://exercism.org/tracks/go/exercises/isbn-verifier/
package isbn

import "strconv"

// IsValidISBN determines if the input string is a valid ISBN-10 or not
func IsValidISBN(isbn string) bool {
	var multiplier int = 10
	var resultant int
	var noOfDigits int

	for i := 0; i < len(isbn); i++ {
		// 'X' denotes a 10
		if isbn[i] == 'X' && multiplier == 1 {
			resultant += 10 * multiplier
			multiplier--
			noOfDigits++
		} else {
			// Determine if the current letter is a digit
			if newInt, err := strconv.Atoi(string(isbn[i])); err == nil {
				resultant += newInt * multiplier
				multiplier--
				noOfDigits++
			}
		}
	}
	if resultant%11 == 0 && noOfDigits == 10 {
		return true
	} else {
		return false
	}
}
