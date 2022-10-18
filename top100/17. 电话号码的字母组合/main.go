/**
* @Author: Chao
* @Date: 2022/8/19 21:34
* @Version: 1.0
 */

package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("23"))
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	var dict = map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	res := dict[digits[0]]
	for i := 1; i < len(digits); i++ {
		tmp := dict[digits[i]]
		list := make([]string, 0)
		for j := 0; j < len(tmp); j++ {
			for k := 0; k < len(res); k++ {
				list = append(list, res[k]+tmp[j])
			}
		}
		res = list
	}

	return res
}
