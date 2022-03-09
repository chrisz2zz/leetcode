package main

import "fmt"

func main() {
	nums := []int{1,3,1,2,0,5}
	fmt.Println(maxSlidingWindow(nums, 3))
}

func maxSlidingWindow2(nums []int, k int) []int {
	//判断直接返回
	if len(nums) == 0 || len(nums) == 1 {
		return nums
	}

	//结果集
	res := make([]int, 0)

	//双向队列
	stack := make([]int, 0)

	//把头k个元素放进去
	for i := 0; i < k; i++ {

		//判断队尾元素是否小于当前元素,小于直接扔掉
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			stack = stack[:len(stack)-1]
		}

		//当前元素入队
		stack = append(stack, i)
	}

	//把前k个最大放入结果集
	res = append(res, nums[stack[0]])

	for i := k; i < len(nums); i++ {
		//判断窗口位置 移除窗口外的元素
		if i - stack[0] >= k {
			stack = stack[1:]
		}

		//判断队尾元素是否小于当前元素
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			stack = stack[:len(stack)-1]
		}

		//当前元素入队
		stack = append(stack, i)

		//当前队列最大元素加入结果集
		res = append(res, nums[stack[0]])
	}

	return res
}


func maxSlidingWindow(nums []int, k int) []int {
	dq := make(dequeue, 0)

	res := make([]int, 0)

	for i := 0; i < k; i++ {
		for dq.len() > 0 && nums[i] > nums[dq.end()] {
			dq.pop()
		}
		dq.push(i)
	}

	res = append(res, nums[dq.front()])

	for i := k; i < len(nums); i++ {
		if i - dq.front() >= k {
			dq.popFront()
		}

		for dq.len() > 0 && nums[i] > nums[dq.end()] {
			dq.pop()
		}
		dq.push(i)

		res = append(res, nums[dq.front()])
	}
	return res
}

type dequeue []int

func (d *dequeue) push(n int) {
	*d = append(*d, n)
}

func (d *dequeue) pop() int {
	t := (*d)[len(*d)-1]
	*d = (*d)[:len(*d)-1]
	return t
}

func (d *dequeue) front() int {
	return (*d)[0]
}

func (d *dequeue) popFront() int {
	t := (*d)[0]
	*d = (*d)[1:]
	return t
}

func (d *dequeue) end() int {
	return (*d)[len(*d)-1]
}

func (d *dequeue) len() int {
	return len(*d)
}