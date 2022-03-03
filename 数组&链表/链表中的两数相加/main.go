/**
* @Author: Chao
* @Date: 2022/3/3 17:40
* @Version: 1.0
 */

package main

import (
	"fmt"
)

func main() {
	h := &ListNode{
		Val:  7,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  4,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}

	h2 := &ListNode{
		Val:  5,
		Next: &ListNode{
			Val:  6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	r := addTwoNumbers(h, h2)
	for r != nil {
		fmt.Println(r.Val)
		r = r.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}
/*
翻转链表解法
 */
//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	l1 = reserve(l1)
//	l2 = reserve(l2)
//
//	var rs *ListNode
//	t := &ListNode{}
//	carry := 0
//	for {
//		a, b := 0, 0
//		if l1 == nil && l2 == nil {
//			break
//		}
//
//		if l1 != nil {
//			a = l1.Val
//			l1 = l1.Next
//		}
//
//		if l2 != nil {
//			b = l2.Val
//			l2 = l2.Next
//		}
//
//		sum := a+b+carry
//		carry = (sum)/10
//		num := sum % 10
//		if rs == nil {
//			rs = &ListNode{Val: num}
//			t = rs
//		} else {
//			t.Next = &ListNode{Val: num}
//			t = t.Next
//		}
//	}
//	if carry > 0 {
//		t.Next = &ListNode{Val: carry}
//	}
//
//	return reserve(rs)
//}
//
//func reserve(n *ListNode) *ListNode {
//	var prev *ListNode
//	for n != nil {
//		//tmp := n.Next
//		n.Next, prev, n = prev, n, n.Next
//		//prev = n
//		//n = tmp
//	}
//
//	return prev
//}

/*
栈解法
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1, stack2 := make([]int, 0), make([]int, 0)

	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}

	prev := &ListNode{}
	carry := 0
	for {
		if len(stack1) == 0 && len(stack2) == 0 {
			break
		}

		a, b := 0, 0

		if len(stack1) != 0 {
			a = stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
		}

		if len(stack2) != 0 {
			b = stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
		}
		sum := a+b+carry
		carry, sum= sum/10, sum%10
		n := &ListNode{Val: sum}
		n.Next = prev.Next
		prev.Next = n
	}

	if carry > 0 {
		prev.Next = &ListNode{Val: carry, Next: prev.Next}
	}

	return prev.Next
}
