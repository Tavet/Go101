package challenge

import (
	"strings"
	"github.com/Tavet/Go101/structure"
)

const (
	BALANCED = "balanced"
	NOT_BALANCED = "not balanced"
)

func CheckBalance(t string) string {
	var stack structure.Stack
	// remove trailing whitespaces
	var brackets = []rune(strings.TrimSpace(t))
	// var stack []string

	for i := 0; i < len(brackets); i++ {
		char :=  string(brackets[i])
		switch char {
		case "(":
			stack.Push(char)
			continue;
		case ")":
			stack.Pop()
			continue;
		}
	}

	if(stack.IsEmpty()) {
		return BALANCED
	} else {
		return NOT_BALANCED
	}
}
