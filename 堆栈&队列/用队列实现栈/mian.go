/**
* @Author: Chao
* @Date: 2022/3/4 13:01
* @Version: 1.0
 */

package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.Top())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Empty())
}

type MyStack struct {
	input  array
	output array
}

func Constructor() MyStack {
	return MyStack{
		input:  make(array, 0),
		output: make(array, 0),
	}
}

func (this *MyStack) Push(x int) {
	this.input.push(x)
}

func (this *MyStack) Pop() int {
	for len(this.input) != 1 {
		this.output.push(this.input.pop())
	}

	elem := this.input.pop()

	for len(this.output) != 0 {
		this.input.push(this.output.pop())
	}

	return elem
}

func (this *MyStack) Top() int {
	for len(this.input) != 1 {
		this.output.push(this.input.pop())
	}

	elem := this.input.pop()

	for len(this.output) != 0 {
		this.input.push(this.output.pop())
	}

	this.input.push(elem)

	return elem
}

func (this *MyStack) Empty() bool {
	return len(this.input) == 0
}

type array []int

func (a *array) pop() int {
	if len(*a) == 0 {
		return -1
	}

	elem := (*a)[0]

	*a = (*a)[1:len(*a)]

	return elem
}

func (a *array) push(x int) {
	*a = append(*a, x)
}
