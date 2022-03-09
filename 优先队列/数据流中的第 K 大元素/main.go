package main

import "fmt"

func main() {
	obj := Constructor(1, []int{})
	fmt.Println(obj.Add(-3))
	fmt.Println(obj.Add(-2))
	fmt.Println(obj.Add(-4))
	fmt.Println(obj.Add(0))
	fmt.Println(obj.Add(4))
}

type KthLargest struct {
	mini miniStack
	k int
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{k: k}
	for _, num := range nums {
		kl.Add(num)
	}

	return kl
}

func (this *KthLargest) Add(val int) int {
	if len(this.mini) < this.k {
		this.mini.add(val)
	} else {
		if this.mini.peek() < val {
			this.mini.poll()
			this.mini.add(val)
		}
	}

	return this.mini.peek()
}

type miniStack []int

func (m *miniStack) peek() int {
	return (*m)[0]
}

func (m *miniStack) poll() {
	if len(*m) >= 1 {
		*m = (*m)[1:len(*m)]
	}
}

func (m *miniStack) add(n int) {
	i := 0
	for i = 0; i < len(*m); i++ {
		if (*m)[i] > n {
			break
		}
	}

	tmp := make([]int, 0)
	tmp = append(tmp,(*m)[:i]...)

	tmp = append(tmp, n)

	*m = append(tmp, (*m)[i:]...)
}