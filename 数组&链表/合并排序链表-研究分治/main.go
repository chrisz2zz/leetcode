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

	//lists = []*ListNode{{Val: 1}}

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

func mergeKLists2(lists []*ListNode) *ListNode {
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

//func mergeKLists(lists []*ListNode) *ListNode {
//	prev := &ListNode{}
//	tmp := prev
//
//	reserve(lists, tmp)
//
//	return prev.Next
//}
//
//func reserve(lists []*ListNode, r *ListNode) {
//	min := math.MaxInt
//	idx := -1
//	nilCount := 0
//	for i := 0; i < len(lists); i++ {
//		if lists[i] != nil {
//			if cmin(lists[i].Val, min) {
//				min = lists[i].Val
//				idx = i
//			}
//		} else {
//			nilCount ++
//		}
//
//		//if len(lists) - nilCount == 1 {
//		//	if lists[i+1] != nil {
//		//		r.Next = lists[i+1]
//		//	}
//		//	return
//		//}
//	}
//
//	if idx == -1 {
//		return
//	}
//
//	r.Next = &ListNode{Val: min}
//	r = r.Next
//	lists[idx] = lists[idx].Next
//
//	reserve(lists, r)
//}
//
//func cmin(x, y int) bool {
//	return x <= y
//}

func mergeKLists(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode,left,right int) *ListNode {
	if left>right {
		return nil
	}
	if left == right {
		return lists[left]
	}
	mid := left + (right-left)/2
	l := merge(lists,left,mid)
	r := merge(lists,mid+1,right)

	return mergeTwo(l, r)
}

func mergeTwo(l1, l2 *ListNode) *ListNode {
	prev := &ListNode{}
	tmp := prev

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tmp.Next = l1
			l1 = l1.Next
		} else {
			tmp.Next = l2
			l2 = l2.Next
		}
		tmp = tmp.Next
	}

	if l1 != nil {
		tmp.Next = l1
	} else {
		tmp.Next = l2
	}

	return prev.Next
}