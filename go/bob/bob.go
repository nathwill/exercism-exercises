// Package bob encapsulates sullen teen behavior
package bob

import (
	"regexp"
	"strings"
)

const testVersion = 3

// Hey responds with eloquence and charm
func Hey(s string) string {
	// shouted at
	hasLetters, _ := regexp.MatchString("[a-zA-Z]", s)
	if s == strings.ToUpper(s) && hasLetters {
		return "Whoa, chill out!"
	}

	// questioned
	endsWithQuestion, _ := regexp.MatchString("\\?[[:space:]]*$", s)
	if endsWithQuestion {
		return "Sure."
	}

	// silence
	isWhitespace, _ := regexp.MatchString("^[[:space:]]+$", s)
	if len(s) == 0 || isWhitespace {
		return "Fine. Be that way!"
	}

	//normal response
	return "Whatever."
}
