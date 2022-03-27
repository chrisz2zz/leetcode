package main

import "fmt"

func main() {
	nums := []int{5,2,3,3,3,1}
	sort(nums)
	fmt.Println(nums)
}

func sort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, l, h int) {
	if l >= h {
		return
	}
	index := findIndex(nums, l, h)
	quickSort(nums, l, index-1)
	quickSort(nums, index+1, h)
}

func findIndex(nums []int, i, j int) int {
	p := i
	for j > i {
		for j > i && nums[j] >= nums[p] {
			j--
		}

		for j > i && nums[i] <= nums[p] {
			i++
		}

		nums[i], nums[j] = nums[j], nums[i]
	}

	nums[p], nums[i] = nums[i], nums[p]

	return i
}