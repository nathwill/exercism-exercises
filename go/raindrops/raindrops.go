// Package raindrops contains methods related to raindrops
package raindrops

import "fmt"

const testVersion = 3

// Convert converts an integer to a raindrop
func Convert(i int) string {
	var out string

	// divisible by 3, 5, 7?
	if i%3 == 0 {
		out += "Pling"
	}
	if i%5 == 0 {
		out += "Plang"
	}
	if i%7 == 0 {
		out += "Plong"
	}

	// send result
	if len(out) == 0 {
		return fmt.Sprintf("%d", i)
	}

	return fmt.Sprintf("%s", out)
}
