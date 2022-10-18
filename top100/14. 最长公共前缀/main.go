/**
* @Author: Chao
* @Date: 2022/8/19 19:50
* @Version: 1.0
 */

package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	idx := 0

	template := strs[0]
	for _, str := range strs {
		if len(str) < len(template) {
			template = str
		}
	}

	for i := 0; i < len(template); i++ {
		equal := 0
		for _, str := range strs {
			if str[i] == template[i] {
				equal++
				continue
			}
			break
		}

		if equal == len(strs) {
			idx++
			continue
		}
		break
	}

	return template[:idx]
}
