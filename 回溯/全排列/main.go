/**
* @Author: Chao
* @Date: 2022/4/25 10:03
* @Version: 1.0
 */

package main

import "fmt"

func main() {
	//fmt.Println(permute([]int{0, 1}))
	//fmt.Println(permute([]int{1, 2, 3}))
	fmt.Println(permute([]int{5, 4, 6, 2}))
}

var all [][]int

func permute(nums []int) [][]int {
	record := make([]int, 0)

	all = make([][]int, 0)

	res(nums, record)

	return all
}

func res(nums, record []int) {
	if len(nums) == len(record) {
		tmp := make([]int, len(record))
		copy(tmp, record)
		all = append(all, tmp)
		return
	}
	for _, v := range nums {
		if in(v, record) {
			continue
		}

		record = append(record, v)

		res(nums, record)

		record = record[:len(record)-1]
	}
}

func in(v int, record []int) bool {
	for _, r := range record {
		if v == r {
			return true
		}
	}
	return false
}
