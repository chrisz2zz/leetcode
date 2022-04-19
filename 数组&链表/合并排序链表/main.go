package main

import (
	"fmt"
	"sort"
)

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	l3 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  6,
			Next: nil,
		},
	}

	var lists []*ListNode
	lists = append(lists, l1)
	lists = append(lists, l2)
	lists = append(lists, l3)

	m := mergeKLists(lists)

	for m != nil {
		fmt.Println(m.Val)
		m = m.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	vals := make([]int, 0)
	for i := 0; i < len(lists); i++ {
		for lists[i] != nil  {
			vals = append(vals, lists[i].Val)
			lists[i] = lists[i].Next
		}
	}

	sort.Ints(vals)

	head := &ListNode{}

	tmp := head

	for i := 0; i < len(vals); i++ {
		tmp.Next = &ListNode{Val: vals[i]}
		tmp = tmp.Next
	}

	return head.Next
}