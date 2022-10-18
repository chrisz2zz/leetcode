/**
* @Author: Chao
* @Date: 2022/8/9 15:48
* @Version: 1.0
 */

package main

import "fmt"

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
}

func findAnagrams(s string, p string) []int {
	pm := make(map[byte]int)
	sm := make(map[byte]int)

	for i := 0; i < len(p); i++ {
		pm[p[i]]++
	}
	res := make([]int, 0)
	l, r := 0, 0
	valid := 0
	for r < len(s) {
		c := s[r]
		r++
		if _, ok := pm[c]; ok {
			sm[c]++
			if sm[c] == pm[c] {
				valid++
			}
		}

		if r-l >= len(p) {
			if valid == len(pm) {
				res = append(res, l)
			}
			d := s[l]
			l++
			if _, ok := pm[d]; ok {
				if pm[d] == sm[d] {
					valid--
				}
				sm[d]--
			}
		}
	}

	return res
}
