package challenge

import "strings"

const (
	BALANCED = "balanced"
	NOT_BALANCED = "not balanced"
)
func CheckBalance(t string) string {
	// remove trailing whitespaces
	var brackets = []rune(strings.TrimSpace(t))
	var stack string = []

	for i := 0; i < len(brackets); i++ {
		char :=  string(brackets[i])
		
	}

	return BALANCED
}