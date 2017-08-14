// Package pangram contains methods for handling pangrams
package pangram

import "strings"

const testVersion = 1

// lower case ascii alphabet
const alphabet = "abcdefghijklmnopqrstuvwxyz"

// IsPangram detects if given string is ascii pangram
func IsPangram(s string) bool {
	// detection is case-insensitive
	s = strings.ToLower(s)

	count := 0

	for _, char := range alphabet {
		if strings.Contains(s, string(char)) {
			count += 1
		}
	}

	return count == len(alphabet)
}
