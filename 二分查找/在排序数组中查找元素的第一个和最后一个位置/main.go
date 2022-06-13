package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{2, 2}, 3))
}

func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] >= target {
			r = mid-1
		} else {
			l = mid+1
		}
	}

	if r+1 >= len(nums) || nums[r+1] != target {
		return []int{-1, -1}	
	}

	j := r+1
	for ; j < len(nums); j++ {
		if nums[j] != target {
			break
		}
	}

	return []int{r+1, j-1}
}