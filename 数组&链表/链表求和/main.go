package main

import (
	"fmt"
)

func main() {
	res := addTwoNumbers(&ListNode{
		Val:  6,
		Next: &ListNode{
			Val:  1,
			Next: &ListNode{
				Val:  7,
				Next: nil,
			},
		},
	}, &ListNode{
		Val:  2,
		Next: &ListNode{
			Val:  9,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	})

	//res := addTwoNumbers(&ListNode{
	//	Val:  1,
	//	Next: &ListNode{
	//		Val:  2,
	//		Next: &ListNode{
	//			Val:  3,
	//			Next: nil,
	//		},
	//	},
	//}, &ListNode{
	//	Val:  2,
	//	Next: nil,
	//})

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
	prev := &ListNode{}
	tmp := prev

	add := 0
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + add
		if sum >= 10 {
			add = 1
			sum %= 10
		} else {
			add = 0
		}

		l1 = l1.Next
		l2 = l2.Next

		tmp.Next = &ListNode{Val: sum}
		tmp = tmp.Next
	}


	for l1 != nil {
		sum := l1.Val + add
		if sum >= 10 {
			add = 1
			sum %= 10
		} else {
			add = 0
		}

		l1 = l1.Next

		tmp.Next = &ListNode{Val: sum}
		tmp = tmp.Next
	}

	for l2 != nil {
		sum := l2.Val + add
		if sum >= 10 {
			add = 1
			sum %= 10
		} else {
			add = 0
		}

		l2 = l2.Next

		tmp.Next = &ListNode{Val: sum}
		tmp = tmp.Next
	}

	if add != 0 {
		tmp.Next = &ListNode{Val: add}
	}

	return prev.Next
}