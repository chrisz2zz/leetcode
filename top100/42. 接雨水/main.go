/**
* @Author: Chao
* @Date: 2022/9/12 19:10
* @Version: 1.0
 */

package main

import "fmt"

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func trap(height []int) int {
	l, r := 0, len(height)-1
	lmax, rmax := 0, 0
	ans := 0
	for l <= r {
		if height[l] > lmax {
			lmax = height[l]
		}

		if height[r] > rmax {
			rmax = height[r]
		}

		p := min(lmax, rmax)

		if height[l] <= p {
			ans += p - height[l]
			l++
			continue
		}

		if height[r] <= p {
			ans += p - height[r]
			r--
			continue
		}
	}

	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
