package main

import "fmt"

/*
两个栈:一个是输入栈(input),一个是输出栈(output)

push操作:
	往input里push数据
pop操作:
	从input中pop所以数据到output, 再将output pop一次即为队列的头
	再将output pop到input

	优化点:判断output是否为空,不为空直接调用output pop, 为空时再把input pop的数据push到output
peek操作:
	从input中pop所以数据到output, 将output 末尾元素取出 但不pop删除
	再将output pop到input

	优化点:判断output是否为空,不为空直接调用output 末尾元素, 为空时再把input pop的数据push到output
empty操作:
	直接判断input是否为空

	优化点:判断input和output是否都为空
 */

func main() {
	queue := Constructor()

	queue.Push(1)
	queue.Push(2)
	fmt.Println(queue.Peek())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Empty())
}

type stack []int

func (s *stack) pop() int {
	if len(*s) == 0 {
		return -1
	}

	t := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return t
}

func (s *stack) push(r int) {
	*s = append(*s, r)
}

type MyQueue struct {
	input  stack
	output stack
}

func Constructor() MyQueue {
	return MyQueue{
		input:  make(stack, 0),
		output: make(stack, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.output = make(stack, 0)
	this.input.push(x)
}

func (this *MyQueue) Pop() int {
	for len(this.input) != 0 {
		this.output.push(this.input.pop())
	}

	elem := this.output.pop()

	for len(this.output) != 0 {
		this.input.push(this.output.pop())
	}

	return elem
}

func (this *MyQueue) Peek() int {
	this.output = make(stack, 0)
	for len(this.input) != 0 {
		this.output.push(this.input.pop())
	}

	elem := this.output[len(this.output) - 1]

	for len(this.output) != 0 {
		this.input.push(this.output.pop())
	}
	return elem
}

func (this *MyQueue) Empty() bool {
	return len(this.input) == 0
}