package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println(findMinDifference([]string{"00:00","23:59","00:00"}))
}

func findMinDifference(timePoints []string) int {
	ts := make([]int, len(timePoints))
	for i := 0; i < len(ts); i++ {
		ts[i] = toMin(timePoints[i])
	}

	quickSort(ts, 0, len(ts)-1)

	min, pre := math.MaxInt, 0

	for i := 1; i < len(ts); i++ {
		tmp := acMin(ts[pre], ts[i])
		if tmp < min {
			min = tmp
		}

		pre = i
	}

	t := acMin(ts[0], ts[len(ts)-1])
	if t < min {
		min = t
	}

	return min
}

func toMin(str string) int {
	t, _ := time.Parse("15:04", str)

	return t.Hour()*60 + t.Minute()
}

func acMin(a, b int) int {
	min := math.MaxInt

	if b - a < min {
		min = b-a
	}

	if int(math.Abs(float64(1440 - b + a))) < min {
		min = int(math.Abs(float64(1440 - b + a)))
	}

	return min
}

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}

	idx := findIndx(nums, l, r)
	quickSort(nums, l, idx-1)
	quickSort(nums, idx+1, r)
}

func findIndx(nums []int, l, r int) int {
	idx := l

	for l < r {
		for l < r && nums[r] >= nums[idx] {
			r--
		}

		for l < r && nums[l] <= nums[idx] {
			l++
		}

		nums[l], nums[r] = nums[r], nums[l]
	}
	nums[idx], nums[l] = nums[l], nums[idx]

	return l
}
