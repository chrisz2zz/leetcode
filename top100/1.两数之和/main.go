/**
* @Author: Chao
* @Date: 2022/8/8 20:24
* @Version: 1.0
 */

package main

import (
	"fmt"
)

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}

func twoSum(nums []int, target int) []int {
	index := make(map[int]int)
	for i, num := range nums {
		idx := target - num
		if v, ok := index[idx]; ok {
			return []int{v, i}
		}
		index[num] = i
	}
	return nil
}
