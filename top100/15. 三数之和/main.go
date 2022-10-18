/**
* @Author: Chao
* @Date: 2022/8/21 22:50
* @Version: 1.0
 */

package main

import (
	"fmt"
)

func main() {
	fmt.Println(threeSum([]int{2, 2, 2, 2, 2}))
}

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	quickSort(nums, 0, len(nums)-1)

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			if nums[i]+nums[l]+nums[r] > 0 {
				r--
			} else if nums[i]+nums[l]+nums[r] < 0 {
				l++
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				n1, n2 := l, r
				for l < r && nums[l] == nums[n1] {
					l++
				}
				for l < r && nums[r] == nums[n2] {
					r--
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

	list[r], list[idx] = list[idx], list[r]
	return r
}
