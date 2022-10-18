/**
* @Author: Chao
* @Date: 2022/10/18 20:37
* @Version: 1.0
 */

package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverseLeftWords("abcdefg", 2))
}

func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}
