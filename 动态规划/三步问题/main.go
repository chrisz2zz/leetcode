package main

import (
	"fmt"
)

func main() {
	fmt.Println(waysToStep(1000000))
}

func waysToStep(n int) int {
	dp := make([]int, 3)
	dp[0], dp[1], dp[2] = 1, 2, 4

	if n < 3 {
		return dp[n]
	}

	for i := 3; i < n; i++ {
		t := (dp[2] + dp[1] + dp[0]) % 1000000007
		dp[0] = dp[1]
		dp[1] = dp[2]
		dp[2] = t
	}

	return dp[2]
}
