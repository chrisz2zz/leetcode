/**
* @Author: Chao
* @Date: 2022/4/19 13:00
* @Version: 1.0
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(shortestToChar("loveleetcode", 'e'))
}

func shortestToChar(s string, c byte) []int {
	idx := make([]int, len(s))
	for i := 0; i < len(idx); i++ {
		idx[i] = math.MaxInt
	}

	l, r := -1, -1

	for i := 0; i < len(s); i++ {
		if s[i] == c {
			l = i
		}

		if s[len(s)-1-i] == c {
			r = len(s) - 1 - i
		}

		if l != -1 {
			idx[i] = min(int(math.Abs(float64(i-l))), idx[i])
		}

		if r != -1 {
			idx[len(s)-1-i] = min(int(math.Abs(float64(len(s)-1-i-r))), idx[len(s)-1-i])
		}
	}

	return idx
}

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}
