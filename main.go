package main

import (
	"fmt"
	"strings"

	"go-code-coverage-sample/calc"
)

var greet bool

func main() {
	run()
}

func run() string {
	b := &strings.Builder{}

	// conditionally tested
	if greet {
		b.WriteString("Hello, World!")
	}

	// tested here, but not in its own unit tests
	b.WriteString(fmt.Sprintf("2 * 3 = %d", calc.Multiply(2, 3)))

	return b.String()
}
