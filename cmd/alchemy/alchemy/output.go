package alchemy

import (
	"fmt"
	"github.com/logrusorgru/aurora"
)

// Output represents a
type Output struct {
}

// Success reports a successful test run to the console
func (output Output) Success(test TestCase) {
	fmt.Println(aurora.Green(fmt.Sprintf("\t✔ %s", test.Name())))
}

// Failure reports a failed test to the console
func (output Output) Failure(test TestCase, err string) {
	fmt.Println(aurora.Red(fmt.Sprintf("\t✖ %s", test.Name())))
	fmt.Println(aurora.Red(fmt.Sprintf("\t\t- %s", err)))
}
