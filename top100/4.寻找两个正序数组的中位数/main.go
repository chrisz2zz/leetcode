/**
* @Author: Chao
* @Date: 2022/8/9 16:50
* @Version: 1.0
 */

package main

import "fmt"

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 10, 20, 21, 30}, []int{2, 22, 31, 67, 90}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

}

//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//	nums1 = append(nums1, nums2...)
//
//	if len(nums1) == 0 {
//		return 0
//	}
//
//	quickSort(nums1, 0, len(nums1)-1)
//	if len(nums1)%2 == 0 {
//		mid := len(nums1) / 2
//		return (float64(nums1[mid]) + float64(nums1[mid-1])) / 2
//	} else {
//		return float64(nums1[len(nums1)/2])
//	}
//}
//
//func quickSort(list []int, l, r int) {
//	if r <= l {
//		return
//	}
//
//	idx := findIndex(list, l, r)
//	quickSort(list, l, idx-1)
//	quickSort(list, idx+1, r)
//}
//
//func findIndex(list []int, l, r int) int {
//	p := l
//	for l < r {
//		for l < r && list[r] >= list[p] {
//			r--
//		}
//
//		for l < r && list[l] <= list[p] {
//			l++
//		}
//
//		list[l], list[r] = list[r], list[l]
//	}
//	list[r], list[p] = list[p], list[r]
//	return r
//}
