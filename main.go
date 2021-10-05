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
	recursionChallenge()
	highOrderChallenge()
}

func highOrderChallenge() {
	slice := [] int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Even numbers: %d\n", challenge.Filter(slice, challenge.IsEven))
	fmt.Printf("Odd numbers: %d\n", challenge.Filter(slice, challenge.IsOdd))
}

func recursionChallenge() {
	var n uint64 = 10
	fmt.Printf("%d factorial: %d\n", n, challenge.Factorial(n))
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