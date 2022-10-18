/**
* @Author: Chao
* @Date: 2022/8/23 22:43
* @Version: 1.0
 */

package main

import (
	"fmt"
	"sort"
)

func main() {
	h := mergeKLists([]*ListNode{{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}, {
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}, {
		Val: 2,
		Next: &ListNode{
			Val:  6,
			Next: nil,
		},
	}})
	for h != nil {
		fmt.Println(h.Val)
		h = h.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	list := make([]int, 0)
	for _, node := range lists {
		for node != nil {
			list = append(list, node.Val)
			node = node.Next
		}
	}

	sort.Ints(list)
	head := &ListNode{}
	h := head
	for _, i := range list {
		head.Next = &ListNode{
			Val:  i,
			Next: nil,
		}
		head = head.Next
	}

	return h.Next
}
