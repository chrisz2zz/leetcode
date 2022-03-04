/**
* @Author: Chao
* @Date: 2022/3/4 10:32
* @Version: 1.0
 */

package main

import "fmt"

func main() {
	head := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{
					Val:  4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	reorderList(head)

	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
数组方法
 */
func reorderList2(head *ListNode) {
	list := make([]*ListNode, 0)
	t := head
	for t != nil {
		list = append(list, t)
		t = t.Next
	}

	length := len(list)

	t = head
	for i := length - 1; i > length / 2; i-- {
		if t == list[i] {
			return
		}

		tmp := t.Next
		list[i].Next = tmp
		t.Next = list[i]

		t = tmp
	}
	list[length/2].Next = nil
}

/*
寻找中点  链表翻转  合并链表
 */
func reorderList(head *ListNode)  {
	if head == nil {
		return
	}

	mid := midNode(head)
	l1 := head
	l2 := mid.Next
	mid.Next = nil
	l2 = reserve(l2)
	mergeList(l1, l2)
}

func midNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reserve(n *ListNode) *ListNode {
	var prev, cur *ListNode
	cur = n
	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}
	return prev
}

func mergeList(n1, n2 *ListNode) {
	var l1, l2 *ListNode

	for n1 != nil && n2 != nil {
		l1 = n1.Next
		l2 = n2.Next

		n1.Next = n2
		n1 = l1

		n2.Next = n1
		n2 = l2
	}
}