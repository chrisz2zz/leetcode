/**
* @Author: Chao
* @Date: 2022/10/11 20:50
* @Version: 1.0
 */

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("a good   example"))
}

func reverseWords(s string) string {
	s = strings.Trim(s, " ")
	ss := strings.Split(s, " ")
	idx := 0
	for i := 0; i < len(ss); i++ {
		if ss[i] != "" {
			ss[idx] = ss[i]
			idx++
		}
	}

	l, r := 0, idx-1
	for l < r {
		ss[l], ss[r] = ss[r], ss[l]
		l++
		r--
	}

	return strings.Join(ss[:idx], " ")
}
