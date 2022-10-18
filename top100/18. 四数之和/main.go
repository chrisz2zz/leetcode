/**
* @Author: Chao
* @Date: 2022/8/21 23:49
* @Version: 1.0
 */

package main

import (
	"fmt"
)

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}

func fourSum(nums []int, target int) [][]int {
	quickSort(nums, 0, len(nums)-1)
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			l, r := j+1, len(nums)-1
			for l < r {
				if nums[i]+nums[j]+nums[l]+nums[r] == target {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					n1, n2 := l, r
					for l < r && nums[l] == nums[n1] {
						l++
					}
					for l < r && nums[r] == nums[n2] {
						r--
					}
				} else if nums[i]+nums[j]+nums[l]+nums[r] > target {
					r--
				} else {
					l++
				}
			}
		}
	}
	return res
}

func quickSort(list []int, l, r int) {
	if l >= r {
		return
	}
	idx := findIdx(list, l, r)
	quickSort(list, l, idx-1)
	quickSort(list, idx+1, r)
}

func findIdx(list []int, l, r int) int {
	idx := l
	for l < r {
		for l < r && list[r] >= list[idx] {
			r--
		}

		for l < r && list[l] <= list[idx] {
			l++
		}
		list[r], list[l] = list[l], list[r]
	}
	list[r], list[idx] = list[idx], list[r]
	return r
}
