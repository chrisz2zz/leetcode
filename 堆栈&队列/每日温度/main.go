/**
* @Author: Chao
* @Date: 2022/3/7 13:16
* @Version: 1.0
 */

package main

import "fmt"

func main() {
	temperatures := []int{73,74,75,71,69,72,76,73}
	fmt.Println(dailyTemperatures(temperatures))
}

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))

	//单调栈
	s := make(stack, 0)

	for i, temperature := range temperatures {
		if len(s) == 0 {
			s.push(i)
			continue
		}

		//如果栈不为空 且 当前温度大于栈顶温度
		for len(s) > 0 && temperature > temperatures[s[len(s)-1]] {
			//出栈
			index := s.pop()

			//得出天数
			res[index] = i - index
		}

		//把当前日期温度下标入栈
		s.push(i)
	}

	return res
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