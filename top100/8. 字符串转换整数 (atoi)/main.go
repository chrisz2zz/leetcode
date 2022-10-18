/**
* @Author: Chao
* @Date: 2022/8/17 22:26
* @Version: 1.0
 */

package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(myAtoi("9223372036854775808"))
}

func myAtoi(s string) int {
	s = strings.Trim(s, " ")

	if len(s) == 0 {
		return 0
	}
	negative := false
	if s[0] == '-' {
		negative = true
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	sum := int64(0)
	for _, i := range s {
		if int(i)-48 < 0 || int(i)-48 > 9 {
			break
		}

		sum = sum*10 + int64(i) - 48
		if sum > math.MaxInt32+1 {
			sum = math.MaxInt32 + 1
		}
	}

	if negative {
		sum = -sum
		if sum < math.MinInt32 {
			sum = math.MinInt32
		}
	}

	if sum > math.MaxInt32 {
		sum = math.MaxInt32
	}

	return int(sum)
}
