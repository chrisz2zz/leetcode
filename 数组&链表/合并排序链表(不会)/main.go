package main

import "fmt"

func main() {
	l1 := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  4,
			Next: &ListNode{
				Val:  5,
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

	l3 := &ListNode{
		Val:  2,
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
     Val int
     Next *ListNode
 }

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	basic := &ListNode{Next: lists[0]}

	for i := 1; i <len(lists); i++ {
		head := lists[i]
		h := basic.Next
		for head != nil {
			p := head
			for h != nil || h.Next != nil {
				tmpP := h
				tmp := h.Next
				if p.Val >= h.Val && p.Val <= h.Val {
					tmpP.Next = p
					p.Next = tmp
					h = h.Next
					continue
				}

				if tmp.Next == nil {
					tmp.Next = p
					p.Next = nil
				}
			}
			head = head.Next
		}
	}
	return basic.Next
}