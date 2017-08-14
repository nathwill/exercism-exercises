// Package bob encapsulates sullen teen behavior
package bob

import (
	"regexp"
	"strings"
)

const testVersion = 3

// Hey responds with eloquence and charm
func Hey(s string) string {
	s = strings.TrimSpace(s)

	hasLetters, _ := regexp.MatchString("[a-zA-Z]", s)
	if s == strings.ToUpper(s) && hasLetters {
		return "Whoa, chill out!"
	}

	if strings.HasSuffix(s, "?") {
		return "Sure."
	}

	if s == "" {
		return "Fine. Be that way!"
	}

	return "Whatever."
}
