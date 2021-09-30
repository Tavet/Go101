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
	fmt.Printf("Celsius: %f - Fahrenheit: %f", celsius, challenge.ToFahrenheit(celsius))
}
