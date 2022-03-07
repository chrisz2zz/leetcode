/**
* @Author: Chao
* @Date: 2022/3/7 12:53
* @Version: 1.0
 */

package main

import (
	"fmt"
	"strconv"
)

func main() {
	tokens := []string{"4","13","5","/","+"}
	fmt.Println(evalRPN(tokens))
}

func evalRPN(tokens []string) int {
	s := make(stack, 0)

	op := map[string]func(int,int)int{
		"+":add,
		"-": sub,
		"*": multi,
		"/": divi,
	}


	for _, token := range tokens {
		if f, ok := op[token]; ok {
			y := s.pop()
			x := s.pop()

			x = f(x, y)

			s.push(x)
			continue
		}

		n, _ := strconv.Atoi(token)
		s.push(n)
	}

	return s.pop()
}

func multi(x, y int) int {
	return x*y
}

func divi(x, y int) int {
	return x/y
}

func add(x, y int) int {
	return x+y
}

func sub(x, y int) int {
	return x-y
}

type stack []int

func (s *stack) push(n int) {
	*s = append(*s, n)
}

func (s *stack) pop() int {
	if len(*s) == 0 {
		return -1
	}
	
	n := (*s)[len(*s)-1]
	
	*s = (*s)[:len(*s)-1]
	
	return n
}