package main

import (
	"fmt"
)

func main() {
	res := addTwoNumbers(&ListNode{
		Val:  7,
		Next: &ListNode{
			Val:  1,
			Next: &ListNode{
				Val:  6,
				Next: nil,
			},
		},
	}, &ListNode{
		Val:  5,
		Next: &ListNode{
			Val:  9,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	})

	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}

	add := 0
	h := res
	for l1 != nil && l2 != nil {
		t := l1.Val + l2.Val + add

		add = 0

		if t >= 10 {
			add = 1
		}


	}


	return res
}
