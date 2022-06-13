package main

import (
	"fmt"
)

func main() {
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
}

func minCostClimbingStairs(cost []int) int {
	l1, l2 := 0, 0

	for i := 2; i <= len(cost); i++ {
		l1, l2 = l2, min(l2+cost[i-1], l1+cost[i-2])
	}

	return l2
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
