// Package hamming provides methods for determining Hamming Distance
package hamming

import "errors"

const testVersion = 6

// Distance returns the Hamming Distance between 2 strings
func Distance(a, b string) (int, error) {
	// only valid for strings of equal lengths
	if len(a) != len(b) {
		return -1, errors.New("Undefined for unequal sequence lengths")
	}

	// compare strings
        distance := 0
	for idx, char := range a {
		if b[idx] != byte(char) {
			distance++
		}
	}

	// return computed Hamming Distance
	return distance, nil
}
