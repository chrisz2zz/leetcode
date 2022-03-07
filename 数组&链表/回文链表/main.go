/**
* @Author: Chao
* @Date: 2022/3/7 12:33
* @Version: 1.0
 */

package main

import "fmt"

func main() {
	//head := &ListNode{
	//	Val:  1,
	//	Next: &ListNode{
	//		Val:  2,
	//		Next: &ListNode{
	//			Val:  3,
	//			Next: &ListNode{
	//				Val:  3,
	//				Next: &ListNode{
	//					Val:  2,
	//					Next: &ListNode{
	//						Val:  1,
	//						Next: nil,
	//					},
	//				},
	//			},
	//		},
	//	},
	//}

	head := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  0,
			Next: &ListNode{
				Val: 1,
				Next: nil,
			},
		},
	}

	fmt.Println(isPalindrome(head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head.Next == nil {
		return true
	}

	mid := midNode(head)
	//中间点算前半段链表
	//这样,奇数的时候  前半段比后半段多一个节点
	//偶数的时候前后一样长
	//判断的时候只需要判断后半段是否为nil即可
	p := mid.Next

	p2 := reserve(p)

	p3 := p2
	p4 := head
	for p3 != nil {
		if p3.Val != p4.Val {
			return false
		}

		p3 = p3.Next
		p4 = p4.Next
	}

	//还原链表
	t := reserve(p2)

	mid.Next = t

	return true
}

func midNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reserve(node *ListNode) *ListNode {
	var prev, cur *ListNode
	cur = node

	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}
	return prev
}