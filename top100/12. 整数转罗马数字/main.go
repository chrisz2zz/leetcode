/**
* @Author: Chao
* @Date: 2022/8/18 22:14
* @Version: 1.0
 */

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(intToRoman(1994))
}

func intToRoman(num int) string {
	romans := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	nums := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}

	buf := strings.Builder{}
	for num != 0 {
		for i := len(nums) - 1; i >= 0; i-- {
			if num >= nums[i] {
				num -= nums[i]
				buf.WriteString(romans[i])
				break
			}
		}
	}
	return buf.String()
}
