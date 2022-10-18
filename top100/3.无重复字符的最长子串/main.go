/**
* @Author: Chao
* @Date: 2022/8/9 09:05
* @Version: 1.0
 */

package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
	win := make(map[byte]int)
	l, r := 0, 0
	res := 0

	for r < len(s) {
		c := s[r]
		r++
		win[c]++

		for win[c] > 1 {
			d := s[l]
			l++
			if _, ok := win[d]; ok {
				win[d]--
			}
		}
		res = max(res, r-l)
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//func lengthOfLongestSubstring(s string) int {
//	wind := make(map[byte]int)
//	res := 0
//
//	l, r := 0, 0
//	for r < len(s) {
//		c := s[r]
//		r++
//		wind[c]++
//
//		for wind[c] > 1 {
//			d := s[l]
//			l++
//			if _, ok := wind[d]; ok {
//				wind[d]--
//			}
//		}
//
//		res = max(r-l, res)
//	}
//
//	return res
//}
//
//func max(x, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}

//func lengthOfLongestSubstring(s string) int {
//	win := make(map[byte]int)
//	res := 0
//
//	l, r := 0, 0
//	for r < len(s) {
//		c := s[r]
//		r++
//		//put in window
//		win[c]++
//
//		for win[c] > 1 {
//			d := s[l]
//			win[d]--
//			l++
//		}
//		res = max(res, r-l)
//	}
//
//	return res
//}
//
//func max(x, y int) int {
//	if y > x {
//		return y
//	}
//	return x
//}
