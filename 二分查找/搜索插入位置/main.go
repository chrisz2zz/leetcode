package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1,3,5,6}, 7))
}

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	ans := len(nums)

	for l <= r {
		mid := l + (r-l)/2

		if nums[mid] >= target {
			ans = mid
			r = mid - 1
			continue
		}

		if nums[mid] < target {
			l = mid + 1
			continue
		}
	}

	return ans
}