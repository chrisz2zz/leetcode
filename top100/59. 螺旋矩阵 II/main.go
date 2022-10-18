/**
* @Author: Chao
* @Date: 2022/9/15 22:33
* @Version: 1.0
 */

package main

import "fmt"

func main() {
	fmt.Println(generateMatrix(3))
}

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, n)
	}

	startx, starty := 0, 0
	i, j := 0, 0
	count, offset := 1, 1
	for l := 0; l < n/2; l++ {
		for j = starty; j < n-offset; j++ {
			res[startx][j] = count
			count++
		}
		for i = startx; i < n-offset; i++ {
			res[i][j] = count
			count++
		}
		for ; j > starty; j-- {
			res[i][j] = count
			count++
		}
		for ; i > startx; i-- {
			res[i][j] = count
			count++
		}
		startx++
		starty++
		offset++
	}

	if n%2 == 1 {
		res[startx][starty] = count
	}
	return res
}
