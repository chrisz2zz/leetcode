/**
* @Author: Chao
* @Date: 2022/9/1 22:54
* @Version: 1.0
 */

package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{}, 0))
}

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}

	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] >= target {
			if nums[mid] == target {
				res[0] = mid
			}
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	l, r = 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] <= target {
			if nums[mid] == target {
				res[1] = mid
			}
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return res
}
