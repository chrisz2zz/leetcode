/**
* @Author: Chao
* @Date: 2022/8/8 20:07
* @Version: 1.0
 */

package main

import "fmt"

func main() {
	fmt.Println(binarySearch([]int{1, 1, 1, 1, 1, 1, 1, 1}, 1))
}

func binarySearch(array []int, dst int) int {
	l, r := 0, len(array)
	if r == 0 {
		return -1
	}

	mid := 0
	for l < r {
		mid = l + (r-l)/2
		if array[mid] < dst {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return mid - 1
}
