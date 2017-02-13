// Package greeting should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package greeting

import "fmt"

const testVersion = 3

// HelloWorld needs a comment documenting it as package does.
func HelloWorld(name string) string {
	if name == "" {
		name = "World"
	}

	return fmt.Sprintf("Hello, %s!", name)
}
