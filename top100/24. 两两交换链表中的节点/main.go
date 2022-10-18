/**
* @Author: Chao
* @Date: 2022/8/23 22:59
* @Version: 1.0
 */

package main

import "fmt"

func main() {
	h := swapPairs(&ListNode{
		Val:  1,
		Next: nil,
	})
	for h != nil {
		fmt.Println(h.Val)
		h = h.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	pre := &ListNode{Next: head}
	h, e := pre, pre
	for head != nil {
		for i := 0; i < 2; i++ {
			e = e.Next
			if e == nil {
				return pre.Next
			}
		}

		nex := e.Next
		head, e = reserveTwo(head, e)

		h.Next = head
		e.Next = nex

		head = nex
		h = e
	}
	return pre.Next
}

func reserveTwo(h, e *ListNode) (*ListNode, *ListNode) {
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
