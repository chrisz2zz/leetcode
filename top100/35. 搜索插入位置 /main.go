/**
* @Author: Chao
* @Date: 2022/8/31 22:24
* @Version: 1.0
 */

package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
}

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	ans := 0
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] <= target {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	if nums[ans] < target {
		ans++
	}

	return ans
}
