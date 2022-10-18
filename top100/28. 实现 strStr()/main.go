/**
* @Author: Chao
* @Date: 2022/8/28 09:45
* @Version: 1.0
 */

package main

import "fmt"

func main() {
	fmt.Println(strStr("hello", "ll"))
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	idx := -1

	for i := 0; i < len(haystack); i++ {
		count := 0
		for j := 0; j < len(needle); j++ {
			if i+j < len(haystack) && haystack[i+j] == needle[j] {
				count++
				continue
			}
			break
		}
		if count == len(needle) {
			idx = i
			break
		}
	}
	return idx
}
