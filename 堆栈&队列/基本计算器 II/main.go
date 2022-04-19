package main

import (
	"fmt"
)

func main() {
	s := "42"
	fmt.Println(calculate(s))
}

type f func(x,y int32) int32

func add(x,y int32) int32 {
	return x+y
}

func sub(x,y int32) int32 {
	return x-y
}

func mult(x,y int32) int32 {
	return x*y
}

func divi(x,y int32) int32 {
	return x/y
}

var op = map[int32]f {
	43: add,
	45: sub,
	42: mult,
	47: divi,
}

//   3+5/2
func calculate(s string) int {
	stk := stack{}

	preflag := int32(0)

	num := 0
	for _, i := range s {
		if i == ' ' {
			continue
		}

		isDigit := '0' <= i && i <= '9'

		if isDigit {
			if preflag == '-' {
				stk.push(0 - i - '0')
				preflag = 0
			} else if preflag == '/' || preflag == '*' {
				x := stk.pop()
				stk.push(op[preflag](x, i - '0'))
				preflag = 0
			} else {
				stk.push(i - '0')
			}
		}

		if i == '/' || i == '*' || i == '-' {
			preflag = i
			num = 0
		}
	}

	sum := int32(0)
	for _, i := range stk {
		sum+=i
	}

	return int(sum)
}

type stack []int32

func (s *stack) push(num int32) {
	*s = append(*s, num)
}

func (s *stack) pop() int32 {
	if len(*s) == 0 {
		return -1
	}

	tmp := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return tmp
}