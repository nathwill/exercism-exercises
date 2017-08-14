// Package accumulate contains methods related to accumulation
package accumulate

const testVersion = 1

// Accumulate transforms a specified slice of strings by applying the given function
func Accumulate(operands []string, operator func(string) string) []string {
	var out []string

	for _, operand := range operands {
		out = append(out, operator(operand))
	}

	return out
}
