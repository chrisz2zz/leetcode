package main

import "fmt"

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	r := reverseKGroup(head, 2)
	for r != nil {
		fmt.Println(r.Val)
		r = r.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	pre := &ListNode{Next: head}
	//头游标 尾游标
	idx, e := pre, pre
	for head != nil {
		//寻找k个节点
		for i := 0; i < k; i++ {
			e = e.Next
			if e == nil {
				return pre.Next
			}
		}

		//记录下一个开始
		nex := e.Next

		//反转k
		head, e = reserve(head, e)

		//串表
		idx.Next = head
		e.Next = nex

		//移位
		head = nex

		idx = e
	}

	return pre.Next
}

func reserve(h, e *ListNode) (*ListNode, *ListNode) {
	var pre *ListNode
	head := h
	for pre != e {
		tmp := head.Next
		head.Next = pre
		pre = head
		head = tmp
	}

	return e, h
}
