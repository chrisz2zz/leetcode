/**
* @Author: Chao
* @Date: 2022/4/19 12:46
* @Version: 1.0
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
}

func removeDuplicates(nums []int) int {
	count := 0
	prev := math.MinInt
	for i := 0; i < len(nums); i++ {
		if prev != nums[i] {
			nums[count] = nums[i]
			count++
			prev = nums[i]
		}
	}

	nums = nums[:count]

	return count
}
