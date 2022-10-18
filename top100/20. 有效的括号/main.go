/**
* @Author: Chao
* @Date: 2022/8/22 21:04
* @Version: 1.0
 */

package main

import "fmt"

func main() {
	fmt.Println(isValid("()[]{}"))
}

func isValid(s string) bool {
	list := []rune(s)
	dict := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}

	stack := make([]rune, 0)
	for _, r := range list {
		if v, ok := dict[r]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != v {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, r)
		}
	}

	if len(stack) != 0 {
		return false
	}
	return true
}
