/**
* @Author: Chao
* @Date: 2022/4/1 13:36
* @Version: 1.0
 */

package main

import "fmt"

func main() {
	fmt.Println(fib(4))
}

//暴力  栈溢出
func fib2(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	return fib2(n-1) + fib2(n-2)
}

//备忘录 自顶向下
func fib3(n int) int {
	record := make([]int, n+1)

	return reserve2(n, record)
}

func reserve2(n int, record []int) int {
	if n == 0 || n == 1 {
		return n
	}

	if record[n] != 0 {
		return record[n]
	}

	record[n] = reserve2(n-1, record) + reserve2(n-2, record)
	return record[n]
}

//dp table 自底向上
func fib4(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	record := make([]int, n+1)

	record[0], record[1] = 0, 1

	for i := 2; i <= n; i++ {
		record[i] = record[i-1] + record[i-2]
	}

	return record[n]
}

//状态转移,缩减空间复杂度
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	i_1, i_2 := 1, 0

	for i := 2; i <= n; i++ {
		i_1, i_2 = i_1 + i_2, i_1
	}

	return i_1
}