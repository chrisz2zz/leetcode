/**
* @Author: Chao
* @Date: 2022/8/9 09:56
* @Version: 1.0
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}

func minWindow(s string, t string) string {
	win, need := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	l, r := 0, 0
	idx, length := 0, math.MaxInt
	valid := 0
	for r < len(s) {
		c := s[r]
		r++

		if _, ok := need[c]; ok {
			win[c]++
			if win[c] == need[c] {
				valid++
			}
		}

		for valid == len(need) {
			if r-l < length {
				idx = l
				length = r - l
			}
			d := s[l]
			l++
			if _, ok := need[d]; ok {
				if need[d] == win[d] {
					valid--
				}
				win[d]--
			}
		}
	}

	if length != math.MaxInt {
		return s[idx : idx+length]
	}
	return ""
}
