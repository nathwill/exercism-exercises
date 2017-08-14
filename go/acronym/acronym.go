// Package acronym contains functions for working with acronyms
package acronym

import "strings"

const testVersion = 3

// Abbreviate converts a phrase to it's acronym
func Abbreviate(p string) string {
	acr := ""

	// hyphenated words treated as separate
	p = strings.Replace(p, "-", " ", -1)

	// take the capitalized first letter of each word
	for _, w := range strings.Split(p, " ") {
		acr += strings.ToUpper(w[:1])
	}

	return acr
}
