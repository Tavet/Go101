package structure

type Stack []rune

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(item rune) {
	*s = append(*s, item)
}

func (s *Stack) Pop() (rune, bool) {
	if(s.IsEmpty()) {
		return 0, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index] // Remove the element from the stack. Creates a new array from position 0 to index 
		return element, true
	}
}