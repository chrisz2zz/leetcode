/**
* @Author: Chao
* @Date: 2022/8/24 22:25
* @Version: 1.0
 */

package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

func removeDuplicates(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			nums[count] = nums[i]
			count++
			continue
		}

		if nums[i] != nums[i-1] {
			nums[count] = nums[i]
			count++
		}
	}
	return count
}
