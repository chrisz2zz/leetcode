package main

import "fmt"

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9}, 9))
}

func search(nums []int, target int) int {
	lf, rh := 0, len(nums)-1

	for lf <= rh {
		mid := lf + (rh-lf)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			rh = mid - 1
		} else if nums[mid] < target {
			lf = mid + 1
		}
	}

	return -1
}