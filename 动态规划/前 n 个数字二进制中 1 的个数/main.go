package main

import (
	"fmt"
)

func main() {
	fmt.Println(countBits(8))
}

func countBits(n int) []int {
	dp := make([]int, n+1)
	dp[0] = 0

	for i := 1; i <= n; i++ {
		dp[i] = dp[i>>1] + i%2
	}
	return dp
}
