/**
* @Author: Chao
* @Date: 2022/3/4 13:24
* @Version: 1.0
 */

package main

import "fmt"

func main() {
	s := "3[a]2[bc]"
	fmt.Println(decodeString(s))
}

func decodeString(s string) string {
	result := ""
	stk := make(stack, 0)
	index := map[rune]rune{']':'['}
	for _, r := range s {
		stk.push(r)
	}

	for stk.length() != 0 {
		if  {
			
		}
	}
	
	return result
}

type stack []rune

func (s *stack) push(x rune) {
	*s = append(*s, x)
}

func (s *stack) pop() rune {
	elem := (*s)[s.length()-1]
	(*s) = (*s)[:s.length()-1]
	return elem
}

func (s *stack) length() int {
	return len(*s)
}