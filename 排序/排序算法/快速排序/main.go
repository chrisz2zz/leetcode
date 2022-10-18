package main

import "fmt"

func main() {
	nums := []int{5, 2, 3, 3, 3, 1}
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

func quickSort(list []int, l, r int) {
	if r <= l {
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
	list[idx], list[r] = list[r], list[idx]

	return r
}

//func sort(nums []int) {
//	quickSort(nums, 0, len(nums)-1)
//}
//
//func quickSort(nums []int, l, h int) {
//	if l >= h {
//		return
//	}
//	index := findIndex(nums, l, h)
//	quickSort(nums, l, index-1)
//	quickSort(nums, index+1, h)
//}
//
//func findIndex(nums []int, i, j int) int {
//	p := i
//	for j > i {
//		for j > i && nums[j] >= nums[p] {
//			j--
//		}
//
//		for j > i && nums[i] <= nums[p] {
//			i++
//		}
//
//		nums[i], nums[j] = nums[j], nums[i]
//	}
//
//	nums[p], nums[i] = nums[i], nums[p]
//
//	return i
//}
