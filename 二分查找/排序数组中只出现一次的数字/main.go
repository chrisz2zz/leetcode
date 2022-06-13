package main

import "fmt"

func main() {
	fmt.Println(singleNonDuplicate([]int{1,1,2,2,3,4,4,5,5}))
}

func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	sum := 0
	l, r := 0, 1

	for r < len(nums) {
		sum += nums[l]
		if nums[r] == nums[l] {
			sum-=nums[r]
			l = r+1
			r = l+1
		} else {
			l++
			r = l+1
		}
	}

	if l == len(nums) - 1 {
		return nums[l]
	}

	return sum
}