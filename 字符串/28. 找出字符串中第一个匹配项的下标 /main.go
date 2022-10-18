/**
* @Author: Chao
* @Date: 2022/10/18 20:47
* @Version: 1.0
 */

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strStr("sadbutsad", "sad"))
}

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
