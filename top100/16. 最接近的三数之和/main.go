/**
* @Author: Chao
* @Date: 2022/8/21 23:29
* @Version: 1.0
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(threeSumClosest([]int{1, 1, 1, 0}, 100))
}

func threeSumClosest(nums []int, target int) int {
	quickSort(nums, 0, len(nums)-1)
	min := math.MaxFloat64
	res := 0
	for i := 0; i < len(nums); i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			if nums[i]+nums[l]+nums[r] == target {
				return nums[i] + nums[l] + nums[r]
			} else if nums[i]+nums[l]+nums[r] > target {
				if math.Abs(float64(nums[i]+nums[l]+nums[r]-target)) < min {
					min = math.Abs(float64(nums[i] + nums[l] + nums[r] - target))
					res = nums[i] + nums[l] + nums[r]
				}
				r--
			} else {
				if math.Abs(float64(nums[i]+nums[l]+nums[r]-target)) < min {
					min = math.Abs(float64(nums[i] + nums[l] + nums[r] - target))
					res = nums[i] + nums[l] + nums[r]
				}
				l++
			}
		}
	}

	return res
}

func quickSort(list []int, l, r int) {
	if l >= r {
		return
	}
	idx := findIndex(list, l, r)
	quickSort(list, l, idx-1)
	quickSort(list, idx+1, r)
}

func findIndex(list []int, l, r int) int {
	idx := l
	for l < r {
		for l < r && list[r] >= list[idx] {
			r--
		}

		for l < r && list[l] <= list[idx] {
			l++
		}

		list[l], list[r] = list[r], list[l]
	}

	list[idx], list[r] = list[r], list[idx]
	return r
}
