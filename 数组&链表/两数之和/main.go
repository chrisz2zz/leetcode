/**
* @Author: Chao
* @Date: 2022/2/28 12:36
* @Version: 1.0
 */

package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

/*
一次遍历,额外使用一个map记录之前已经遍历过的值和下标
通过target-num[i]来确定  需要的另一个值是什么  然后去map找
若存在 即可返回i和map-v
*/
func twoSum(nums []int, target int) []int {
	tmp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		index := target - nums[i]
		if v, ok := tmp[index]; ok {
			return []int{v, i}
		}
		tmp[nums[i]] = i
	}
	return nil
}
