package main

import "fmt"

func main() {
	s := "([)]"

	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	stack := make(stack, 0)
	dict := map[rune]rune{')':'(',']':'[','}':'{'}

	for _, i := range s {
		if v, ok := dict[i]; ok {
			if v != stack.pop() {
				return false
			}
			continue
		}

		stack.push(i)
	}

	if len(stack) != 0 {
		return false
	}

	return true
}

type stack []rune

func (s *stack) pop() rune {
	if len(*s) == 0 {
		return -1
	}

	t := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return t
}

func (s *stack) push(r rune) {
	*s = append(*s, r)
}