/**
* @Author: Chao
* @Date: 2022/8/17 21:43
* @Version: 1.0
 */

package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("cbbd"))
}

func longestPalindrome(s string) string {
	idx, length := 0, 0
	for i := 0; i < len(s); i++ {
		a := spread(i, i, s)
		b := spread(i, i+1, s)
		if b > a {
			a = b
		}

		if a > length {
			idx = i
			length = a
		}
	}

	idx = idx - (length-1)/2
	return s[idx : idx+length]
}

func spread(l, r int, s string) int {
	for l >= 0 && r < len(s) {
		if s[l] != s[r] {
			break
		}
		l--
		r++
	}
	return r - l - 1
}
