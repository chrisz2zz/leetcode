/**
* @Author: Chao
* @Date: 2022/8/22 20:50
* @Version: 1.0
 */

package main

import "fmt"

func main() {
	//h := removeNthFromEnd(&ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 2,
	//		Next: &ListNode{
	//			Val: 3,
	//			Next: &ListNode{
	//				Val: 4,
	//				Next: &ListNode{
	//					Val:  5,
	//					Next: nil,
	//				},
	//			},
	//		},
	//	},
	//}, 2)
	h := removeNthFromEnd(&ListNode{
		Val:  1,
		Next: nil,
	}, 1)
	for h != nil {
		fmt.Println(h.Val)
		h = h.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	hair := &ListNode{Next: head}
	slow, fast, pre := hair, hair, hair
	for fast != nil {
		fast = fast.Next
		if n != 0 {
			n--
			continue
		}
		pre = slow
		slow = slow.Next
	}

	pre.Next = slow.Next

	return hair.Next
}
