/**
* @Author: Chao
* @Date: 2022/4/6 13:17
* @Version: 1.0
 */

package main

func main() {

}

func MergeSort(array []int, begin int, end int) {
	// 元素大于1进入递归
	if end - begin > 1 {
		//将数组一分为二
		mid := begin + (end - begin + 1) / 2

		//左
		MergeSort(array, begin, mid)
		//右
	}
}