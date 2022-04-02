/**
* @Author: Chao
* @Date: 2022/4/2 11:09
* @Version: 1.0
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
}

//超时
// 暴力
// 假设目标金额为 n，给定的硬币个数为 k，
// 那么递归树最坏情况下高度为 n（全用面额为 1 的硬币），
// 然后再假设这是一棵满 k 叉树，则节点的总数就是 k^(n-1)。
//接下来看每个子问题的复杂度，由于每次递归包含一个 for 循环，
//复杂度为 O(k)，相乘得到总时间复杂度为 O(k^n)，指数级别。
func coinChange2(coins []int, amount int) int {
	//base case判断
	//即 amount=0表示凑出正好的零钱
	// amount < 0表示没有凑出正好的零钱
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}

	result := math.MaxInt
	for _, coin := range coins {
		sub := coinChange(coins, amount-coin)
		//说明该次递归没有凑出正好的零钱,直接跳过
		if sub == -1 {
			continue
		}
		//走到这说明,叶子结点为0,正好凑出零钱
		//根据下层的返回,取出该层递归,所有元素的子问题最小值
		if result > sub+1 {
			result = sub + 1
		}
	}

	if result != math.MaxInt {
		return result
	}

	return -1
}

// 增加记录,减少重复计算
// 子问题不会超过amount的大小O(n)
// 每层递归执行len(coins) O(k)
// 所有时间复杂度 O(kn)
func coinChange(coins []int, amount int) int {
	records := make([]int, amount+1)
	for i := 0; i < len(records); i++ {
		records[i] = math.MinInt
	}

	return reserve(coins, records, amount)
}

func reserve(coins, records []int, amount int) int {
	//base case判断
	//即 amount=0表示凑出正好的零钱
	// amount < 0表示没有凑出正好的零钱
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}

	if records[amount] != math.MinInt {
		return records[amount]
	}

	result := math.MaxInt
	for _, coin := range coins {
		sub := reserve(coins, records, amount-coin)
		//说明该次递归没有凑出正好的零钱,直接跳过
		if sub == -1 {
			continue
		}
		//走到这说明,叶子结点为0,正好凑出零钱
		//根据下层的返回,取出该层递归,所有元素的子问题最小值
		if result > sub+1 {
			result = sub + 1
		}
	}

	records[amount] = -1
	if result != math.MaxInt {
		records[amount] = result
	}

	return records[amount]
}
