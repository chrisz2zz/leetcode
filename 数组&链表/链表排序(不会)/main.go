package main

import "fmt"

func main() {
	head := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}

	s := sortList(head)

	for s != nil {
		fmt.Println(s.Val)
		s = s.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	list := make([]*ListNode, 0)
	for head != nil {
		list = append(list, head)
		head = head.Next
	}

	sort(list)

	prev := &ListNode{}
	t := prev
	for _, node := range list {
		t.Next = node
		t = t.Next
	}
	t.Next = nil
	return prev.Next
}

func sort(list []*ListNode) {
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if list[j].Val < list[i].Val {
				list[j], list[i] = list[i], list[j]
			}
		}
	}
}
