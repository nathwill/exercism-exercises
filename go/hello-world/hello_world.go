package hello

import (
	"fmt"
)

const testVersion = 2

// HelloWorld greets user by name if given, or greets everyone
func HelloWorld(n string) string {
	greet := "World"

	if len(n) > 0 {
		greet = n
	}

	return fmt.Sprintf("Hello, %s!", greet)
}
