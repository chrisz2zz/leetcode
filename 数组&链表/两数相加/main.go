/**
* @Author: Chao
* @Date: 2022/2/28 13:07
* @Version: 1.0
 */

package main

import "fmt"

func main() {
	t := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}

	ttt := reserve(t)
	for ttt != nil {
		fmt.Println(ttt.Val)
		ttt = ttt.Next
	}
	return
	l1 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val:  9,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}

	l2 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val:  9,
				Next: nil,
			},
		},
	}

	r := addTwoNumbers(l1, l2)
	for r != nil {
		fmt.Println(r.Val)
		r = r.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var rs *ListNode
	t := &ListNode{}
	carry := 0
	for l1 != nil || l2 != nil {
		a, b := 0, 0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}

		sum := a + b + carry
		sum, carry = sum%10, sum/10
		if rs == nil {
			rs = &ListNode{Val: sum}
			t = rs
		} else {
			t.Next = &ListNode{Val: sum}
			t = t.Next
		}
	}
	if carry > 0 {
		t.Next = &ListNode{Val: carry}
	}
	return rs
}

func reserve(node *ListNode) *ListNode {
	var cur *ListNode
	prev := &ListNode{Next: node}
	cur = node
	for cur != nil && cur.Next != nil {
		tmp := cur.Next
		tmp2 := prev.Next
		prev.Next = tmp
		cur.Next = tmp.Next
		tmp.Next = tmp2
	}
	return prev.Next
}
