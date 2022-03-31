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
	if head == nil || head.Next == nil {
		return head
	}

	list := make([]*ListNode, 0)
	for head != nil {
		list = append(list, head)
		head = head.Next
	}

	sort(list, 0, len(list)-1)

	prev := &ListNode{}
	t := prev
	for _, node := range list {
		t.Next = node
		t = t.Next
	}
	t.Next = nil
	return prev.Next
}

func sort(list []*ListNode, l, h int) {
	if l >= h {
		return
	}

	index := findIndex(list, l, h)
	sort(list, l, index-1)
	sort(list, index+1, h)
}

func findIndex(list []*ListNode, l, h int) int {
	if l >= h {
		return l
	}

	tmp := l

	for h > l {
		for h > l && list[h].Val >= list[tmp].Val {
			h--
		}

		for h > l && list[l].Val <= list[tmp].Val {
			l++
		}

		list[l], list[h] = list[h], list[l]
	}
	list[tmp], list[l] = list[l], list[tmp]

	return l
}