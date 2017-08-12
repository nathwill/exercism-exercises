package hello

import (
	"fmt"
)

const testVersion = 2

// greet user by name if given, or greet everyone
func HelloWorld(n string) string {
	greet := "World"

	if len(n) > 0 {
		greet = n
	}

	return fmt.Sprintf("Hello, %s!", greet)
}
