/**
* @Author: Chao
* @Date: 2022/8/24 22:16
* @Version: 1.0
 */

package main

import "fmt"

func main() {
	h := reverseKGroup(&ListNode{
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
	}, 2)
	for h != nil {
		fmt.Println(h.Val)
		h = h.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	pre := &ListNode{Next: head}
	h, e := pre, pre
	for head != nil {
		for i := 0; i < k; i++ {
			e = e.Next
			if e == nil {
				return pre.Next
			}
		}

		nex := e.Next

		head, e = reserve(head, e)

		h.Next = head
		e.Next = nex

		h = e
		head = nex
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
