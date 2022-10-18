/**
* @Author: Chao
* @Date: 2022/9/1 23:01
* @Version: 1.0
 */

package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestValidParentheses("()(())"))
}

func longestValidParentheses(s string) int {
	if len(s) < 2 {
		return 0
	}

	max := 0

	l, r := 0, 2
	for {
		if r > len(s) {
			return max
		}

		tmp := judge(s[l:r])
		if tmp > max {
			max = tmp
			r += 2
		} else {
			l++
			r++
		}
	}
}

func judge(s string) int {
	length := 0

	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return length
			}

			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if b == '(' {
				length += 2
			} else {
				return length
			}
		}
	}
	return length
}
