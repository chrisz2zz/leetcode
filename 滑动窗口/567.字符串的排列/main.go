/**
* @Author: Chao
* @Date: 2022/8/9 16:25
* @Version: 1.0
 */

package main

import "fmt"

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
}

func checkInclusion(s1 string, s2 string) bool {
	wind, need := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}

	l, r := 0, 0
	vaild := 0
	for r < len(s2) {
		c := s2[r]
		r++
		if _, ok := need[c]; ok {
			wind[c]++
			if wind[c] == need[c] {
				vaild++
			}
		}

		if r-l >= len(s1) {
			if vaild == len(need) {
				return true
			}

			d := s2[l]
			l++
			if _, ok := need[d]; ok {
				if need[d] == wind[d] {
					vaild--
				}
				wind[d]--
			}
		}
	}
	return false
}
