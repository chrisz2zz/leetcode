/**
* @Author: Chao
* @Date: 2022/8/18 21:44
* @Version: 1.0
 */

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(convert("AB", 1))
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	length := len(s)
	res := make([][]string, numRows)
	for i := 0; i < numRows; i++ {
		res[i] = make([]string, length)
	}

	l, h := 0, 0
	flag := false
	for _, char := range s {
		res[h][l] = string(char)
		if h == numRows-1 {
			flag = true
		}
		if h == 0 {
			flag = false
		}

		if !flag {
			h++
		} else {
			l++
			h--
		}
	}

	buf := strings.Builder{}
	for _, re := range res {
		buf.WriteString(strings.Join(re, ""))
	}

	return buf.String()
}
