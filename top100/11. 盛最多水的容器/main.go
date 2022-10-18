/**
* @Author: Chao
* @Date: 2022/8/18 22:03
* @Version: 1.0
 */

package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 1}))
}

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	area := 0

	for r != l {
		h := min(height[l], height[r])
		area = max(area, h*(r-l))
		if h == height[l] {
			l++
			continue
		}
		r--
	}
	return area
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
