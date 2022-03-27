package main

import "fmt"

func main() {
	l1 := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	l := mergeTwoLists(l1, l2)

	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	prev := head

	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			prev.Next = list2
			list2 = list2.Next
		} else {
			prev.Next = list1
			list1 = list1.Next
		}
		prev = prev.Next
	}

	if list1 != nil {
		prev.Next = list1
	} else {
		prev.Next = list2
	}

	return head.Next
}
