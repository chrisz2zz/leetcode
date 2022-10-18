/**
* @Author: Chao
* @Date: 2022/9/13 22:02
* @Version: 1.0
 */

package main

import "fmt"

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}

var res [][]int

func combinationSum(candidates []int, target int) [][]int {
	res = make([][]int, 0)
	tmp := make([]int, 0)
	dfs(candidates, tmp, target)
	return res
}

func dfs(candidates, tmp []int, target int) {
	for i := 0; i < len(candidates); i++ {
		if target == 0 {
			l := make([]int, len(tmp))
			copy(l, tmp)
			res = append(res, l)
			return
		} else if target < 0 {
			return
		}

		tmp = append(tmp, candidates[i])

		dfs(candidates[i:], tmp, target-candidates[i])

		tmp = tmp[:len(tmp)-1]
	}
}
