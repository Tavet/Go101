package main

import (
	"fmt"
	"github.com/Tavet/Go101/challenge"
)

// Enumeration
const (
	UNKNOWN = iota
	FEMALE
	_ // Skips one
	MALE
)

func main() {
	// Conevrt celsius to fahrenheint
	var celsius challenge.Celsius = 100
	fmt.Printf("Celsius: %f - Fahrenheit: %f\n", celsius, challenge.ToFahrenheit(celsius))

	// Brackets balance challenge
	var brackets string = "(())()()((()))"
	fmt.Printf("%s is %s\n", brackets, challenge.CheckBalance(&brackets))

	// Season of a Month challenge
	fmt.Printf("November is in %s season\n", challenge.Season(11))
}
