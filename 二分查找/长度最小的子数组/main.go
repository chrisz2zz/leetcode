package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}

func minSubArrayLen(target int, nums []int) int {
	sum, l, r, min := 0, 0, 0, math.MaxInt
	for r < len(nums) {
		sum += nums[r]
		for sum >= target {
			if r-l+1 < min {
				min = r - l + 1
			}
			sum -= nums[l]
			l++
		}
		r++
	}

	if math.MaxInt == min {
		return 0
	}
	return min
}

// func minSubArrayLen(target int, nums []int) int {
// 	min := math.MaxInt
// 	for i := 0; i < len(nums); i++ {
// 		res := judge(target, nums[i:])
// 		if res != 0 {
// 			if res < min {
// 				min = res
// 			}
// 		}
// 	}

// 	if min == math.MaxInt {
// 		return 0
// 	}

// 	return min
// }

// func judge(target int, nums []int) int {
// 	sum := 0
// 	for i := 0; i < len(nums); i++ {
// 		sum += nums[i]
// 		if sum >= target {
// 			return len(nums[:i+1])
// 		}
// 	}

// 	if sum < target {
// 		return 0
// 	}

// 	return len(nums)
// }
