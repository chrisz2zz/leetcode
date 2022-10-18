/**
* @Author: Chao
* @Date: 2022/8/28 08:12
* @Version: 1.0
 */

package main

import "fmt"

func main() {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
}

func removeElement(nums []int, val int) int {
	l := 0
	for _, num := range nums {
		if num != val {
			nums[l] = num
			l++
		}
	}
	return l
}
