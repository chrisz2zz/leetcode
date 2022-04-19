/**
* @Author: Chao
* @Date: 2022/4/18 10:43
* @Version: 1.0
 */

package main

import "fmt"

func HeapSort(arr []int) {
	// 创建堆
	heap := BuildHeap(arr)
	var sortedArr []int
	for len(heap) > 0 {
		sortedArr = append(sortedArr, heap.Pop())
	}

	fmt.Println(sortedArr)
}

func main() {
	//输出 [3 8 10 15 15 16 17 19 24 30 33]
	HeapSort([]int{33, 24, 8, 3, 10, 15, 16, 15, 30, 17, 19})
}


type Heap []int

func (h Heap) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Heap) less(i, j int) bool {
	return h[i] < h[j]
}

func (h Heap) up(i int) {
	for {
		f := (i-1)/2 // 是否需要判断i为负数break
		if i == f || h.less(f, i) {
			break
		}
		h.swap(f, i)
		i = f
	}
}

func (h *Heap) Push(x int) {
	*h = append(*h, x)
	h.up(len(*h)-1)
}

func (h Heap) down(i int) {
	for {
		l := 2*i+1
		r := 2*i+2
		if l >= len(h) {
			break
		}
		j := l
		if r < len(h) && h.less(r, l) {
			j = r
		}

		if h.less(i, j) {
			break
		}

		h.swap(i, j)
		i = j
	}
}

// 删除堆中位置为i的元素
// 返回被删元素的值
func (h *Heap) Remove(i int) (int, bool) {
	n := len(*h)-1

	if i < 0 || i > n {
		return 0, false
	}

	h.swap(i, n) // 用最后的元素值替换被删除元素
	// 删除最后的元素
	x := (*h)[n]
	*h = (*h)[0:n]
	// 如果当前元素大于父节点，向下筛选
	if i == 0 || (*h)[i] > (*h)[(i-1)/2] {
		h.down(i)
	} else {  // 当前元素小于父节点，向上筛选
		h.up(i)
	}

	return x, true
}

func (h *Heap) Pop() int {
	x, _ := h.Remove(0)
	return x
}

func BuildHeap(arr []int) Heap {
	h := Heap(arr)
	n := len(h)
	// 从第一个非叶子节点，到根节点
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i)
	}

	return h
}