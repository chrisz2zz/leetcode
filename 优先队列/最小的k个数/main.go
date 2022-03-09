package main

import (
	"container/heap"
	"fmt"
)

func main() {
	arr := []int{3,2,1}
	fmt.Println(getLeastNumbers(arr, 2))
}

func getLeastNumbers(arr []int, k int) []int {
	h := make(miniheap, 0)
	heap.Init(&h)
	for _, i := range arr {
		heap.Push(&h, i)
	}
	res := make([]int, 0)
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(&h).(int))
	}

	return res
}

type miniheap []int

func (m *miniheap) Len() int {
	return len(*m)
}

func (m miniheap) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m miniheap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *miniheap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *miniheap) Pop() interface{} {
	t := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return t
}
