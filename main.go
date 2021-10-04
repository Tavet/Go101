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
	celsiusChallenge()
	bracketsBalanceChallenge()
	seasonChallenge()
}

func celsiusChallenge() {
	var celsius challenge.Celsius = 100
	fmt.Printf("Celsius: %f - Fahrenheit: %f\n", celsius, challenge.ToFahrenheit(celsius))
}

func bracketsBalanceChallenge() {
	var brackets string = "((test) He llo)()"
	fmt.Printf("%s is %s\n", brackets, challenge.CheckBalance(&brackets))
}

func seasonChallenge() {
	fmt.Printf("November is in %s season\n", challenge.Season(11))
}