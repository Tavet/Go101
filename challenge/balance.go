package challenge

import (
	"strings"
	"github.com/Tavet/Go101/structure"
)

const (
	BALANCED = "balanced"
	NOT_BALANCED = "not balanced"
)

func CheckBalance(t *string) string {
	var stack structure.Stack

	for _, char := range strings.TrimSpace(*t) {
		switch string(char) {
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
