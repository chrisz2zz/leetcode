package main

import "fmt"

func main() {
	nums := []int{5,2,3,3,3,1}
	fmt.Println(sortArray(nums))
}

func sortArray(nums []int) []int {
	sort(nums, 0, len(nums)-1)
	return nums
}

func sort(nums []int, l, h int) {
	if l >= h {
		return
	}

	index := quickSort(nums, l, h)

	sort(nums, l, index - 1)
	sort(nums, index + 1, h)
}

func quickSort(nums []int, l, h int) int {
	index := l
	for l < h {
		for h > l && nums[h] >= nums[index] {
			h--
		}

		for h > l && nums[l] <= nums[index] {
			l++
		}

		nums[l], nums[h] = nums[h], nums[l]
	}

	nums[index], nums[l] = nums[l], nums[index]

	return l
}