/**
* @Author: Chao
* @Date: 2022/8/28 09:54
* @Version: 1.0
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(divide(10, -3))
}

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}

	flag := false
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		flag = true
	}

	dividend = int(math.Abs(float64(dividend)))
	divisor = int(math.Abs(float64(divisor)))

	count := 0
	for dividend >= 0 {
		dividend -= divisor
		count++
	}

	if count == 0 {
		count = 1
	} else {
		count--
	}
	if flag {
		count = 0 - count
	}

	if count > math.MaxInt32 || count < math.MinInt32 {
		count = math.MaxInt32
	}

	return count
}
