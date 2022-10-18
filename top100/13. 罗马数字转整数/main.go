/**
* @Author: Chao
* @Date: 2022/8/21 23:19
* @Version: 1.0
 */

package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	dict := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	ans := 0
	for i := 0; i < len(s); i++ {
		if i+1 <= len(s)-1 && dict[s[i]] < dict[s[i+1]] {
			ans -= dict[s[i]]
		} else {
			ans += dict[s[i]]
		}
	}
	return ans
}
