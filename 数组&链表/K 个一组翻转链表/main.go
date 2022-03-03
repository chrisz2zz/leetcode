/**
* @Author: Chao
* @Date: 2022/3/2 12:04
* @Version: 1.0
 */

package main

import (
	"fmt"
)

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

	//head := &ListNode{
	//	Val:  1,
	//	Next: &ListNode{
	//		Val:  2,
	//		Next: nil,
	//	},
	//}

	//fmt.Println(&head)
	//fmt.Println(unsafe.Pointer(head))

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

/*
分组翻转, 不断形成分组翻转小链表, 并移动分组虚拟头用以连接
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	//虚拟头 也用作返回
	h := &ListNode{Next: head}
	prev := h
	for head != nil {
		tail := prev

		// 确定该次分组是否满足,取出最后一个节点
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return h.Next
			}
		}

		// 取出该分组的下一个节点,作为后续头结点
		nex := tail.Next

		// 翻转该分组链表,返回头尾节点
		head, tail = reserve(head, tail)

		// 该分组的虚拟头结点next指向新的头
		prev.Next = head
		// 该分组的尾结点指向下一个分组的头结点
		tail.Next = nex
		// 挪动虚拟头到下一分组的虚拟头
		prev = tail
		// 挪到头结点到下一分组的真实头
		head = tail.Next
	}

	return h.Next
}

/*
链表翻转
*/
func reserve(head, tail *ListNode) (*ListNode, *ListNode) {
	var p, prev *ListNode
	p = head

	for prev != tail {
		tmp := p.Next
		p.Next = prev
		prev = p
		p = tmp
	}
	return tail, head
}

/*
栈
 */
func reverseKGroup2(head *ListNode, k int) *ListNode {
	//数组作为栈
	stack := make([]*ListNode, 0)
	// 虚拟头
	hair := &ListNode{Next: head}

	prev := hair
	for {
		//入栈操作, k个一组
		for i := 0; i < k; i++ {
			//判断head为空的话
			//把栈中的数据从头到尾按序连成链表
			//然后把链表加到prev后
			if head == nil {
				prev.Next = last(&stack)
				return hair.Next
			}
			//一个个入栈
			push(head, &stack)
			//后移
			head = head.Next
		}

		//出栈
		//从尾部一个个弹出
		//直到栈空  表示该组翻转完成
		for len(stack) != 0 {
			n := pop(&stack)
			prev.Next = n
			prev = prev.Next
		}
	}
}

func push(node *ListNode, list *[]*ListNode)  {
	*list = append(*list, node)
}

func pop(list *[]*ListNode) *ListNode {
	node := (*list)[len(*list) - 1]
	*list = (*list)[:len(*list) - 1]
	return node
}

func last(list *[]*ListNode) *ListNode {
	if len(*list) == 0 {
		return nil
	}

	h := &ListNode{Next: (*list)[0]}
	prev := h.Next
	for i := 1; i < len(*list); i++ {
		prev.Next = (*list)[i]

		prev = prev.Next
	}

	return h.Next
}