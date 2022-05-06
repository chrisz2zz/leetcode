package main

import "fmt"

func main() {
	a := Constructor(3)
	fmt.Println(a.Next(1))
	fmt.Println(a.Next(10))
	fmt.Println(a.Next(3))
	fmt.Println(a.Next(5))
}

type MovingAverage struct {
	data  []int
	size int
	sum   int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{
		data: make([]int, 0, size),
		size: size,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	this.data = append(this.data, val)
	this.sum += val
	if len(this.data) > this.size {
		this.sum -= this.data[0]
		this.data = this.data[len(this.data)-this.size:]
	}

	return float64(this.sum) / float64(len(this.data))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
