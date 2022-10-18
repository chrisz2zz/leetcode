/**
* @Author: Chao
* @Date: 2022/8/22 21:11
* @Version: 1.0
 */

package main

import "fmt"

func main() {
	h := mergeTwoLists(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}, &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	src := &ListNode{}
	h := src
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			src.Next = list1
			list1 = list1.Next
			src = src.Next
			continue
		}

		src.Next = list2
		list2 = list2.Next
		src = src.Next
	}

	if list1 != nil {
		src.Next = list1
	} else if list2 != nil {
		src.Next = list2
	}

	return h.Next
}
