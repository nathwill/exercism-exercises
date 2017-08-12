// Package hamming provides methods for determining Hamming Distance
package hamming

import "errors"

const testVersion = 6

// func Distance returns the Hamming Distance between 2 strings
func Distance(a, b string) (int, error) {
	distance := 0

	// only valid for strings of equal lengths
	if len(a) != len(b) {
		return -1, errors.New("Undefined for unequal sequence lengths")
	}

	// compare strings
	for idx, char := range a {
		if b[idx] != byte(char) {
			distance += 1
		}
	}

	// return computed Hamming Distance
	return distance, nil
}
