/**
* @Author: Chao
* @Date: 2022/8/17 21:37
* @Version: 1.0
 */

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(11))
}

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	i, j := 0, len(str)-1
	for i <= j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}
