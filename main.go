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
	var celsius challenge.Celsius = 100
	fmt.Printf("Celsius: %f - Fahrenheit: %f\n", celsius, challenge.ToFahrenheit(celsius))

	var brackets string = "((())"
	fmt.Printf("%s is %s", brackets, challenge.CheckBalance(brackets))
}
