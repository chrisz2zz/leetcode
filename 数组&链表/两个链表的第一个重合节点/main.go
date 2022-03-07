/**
* @Author: Chao
* @Date: 2022/3/3 17:04
* @Version: 1.0
 */

package main

import "fmt"

func main() {
	//a1 := &ListNode{Val:  4, Next: nil}
	a2 := &ListNode{Val: 1, Next: nil}
	//a1.Next = a2

	b1 := &ListNode{Val: 5, Next: nil}
	b2 := &ListNode{Val: 0, Next: nil}
	b3 := &ListNode{Val: 1, Next: nil}
	b1.Next = b2
	b2.Next = b3

	//c1 := &ListNode{Val:  8, Next: &ListNode{Val:  4, Next: &ListNode{Val:  5, Next: nil}}}
	//a2.Next = c1
	//b3.Next = c1

	fmt.Println(getIntersectionNode(a2, b1))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
延迟移动短的链表
分别遍历两个链表,记录长度
找到长的那一条,开始移动,直到剩余长度与短的那条相等时,移动短的
若相等,则为相交点
不相等,不相交
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	distanceA, distanceB := 0, 0

	for a != nil {
		distanceA++
		a = a.Next
	}

	for b != nil {
		distanceB++
		b = b.Next
	}

	a, b = headA, headB
	if distanceB > distanceA {
		for i := distanceB; i > 0; i-- {
			if i == distanceA {
				distanceA--
				if a == b {
					return b
				}
				a = a.Next
			}
			b = b.Next
		}

		return nil
	}

	for i := distanceA; i > 0; i-- {
		if distanceB == i {
			distanceB--
			if a == b {
				return b
			}
			b = b.Next
		}
		a = a.Next
	}

	return nil
}
