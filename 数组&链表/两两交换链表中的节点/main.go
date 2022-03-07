package main

import "fmt"

//func main() {
//	head := &ListNode{}
//	tmp := head
//	for i := 1; i <= 5; i++ {
//		tmp.Val = i
//
//		if i != 5 && tmp.Next == nil {
//			tmp.Next = &ListNode{}
//			tmp = tmp.Next
//		}
//	}
//
//	rs := swapPairs(head)
//
//	for rs != nil {
//		fmt.Println(rs.Val)
//		rs = rs.Next
//	}
//}
//
//type ListNode struct {
//	Val int
//	Next *ListNode
//}
//
//func swapPairs(head *ListNode) *ListNode {
//	dummy := &ListNode{}
//	dummy.Next = head
//	prev := dummy
//
//	for head != nil && head.Next != nil {
//		s := head.Next
//		head.Next = s.Next
//		s.Next = head
//		prev.Next = s
//
//		prev = head
//		head = head.Next
//	}
//
//	return dummy.Next
//}

func main() {
	h := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}

	r := swapPairs(h)
	for r != nil {
		fmt.Println(r.Val)
		if r.Next == nil {
			break
		}
		r = r.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs2(head *ListNode) *ListNode {
	fake := &ListNode{0, head}
	tmp := fake
	for tmp.Next != nil && tmp.Next.Next != nil {
		a := tmp.Next
		b := tmp.Next.Next

		tmp.Next = b
		a.Next = b.Next
		b.Next = a
		tmp = a
	}
	return fake.Next
}

/*
创造一个虚拟头指向实际头节点
然后将虚拟头指向第二个节点
第一个节点指向第二个节点的Next
第二个节点指向第一个节点
然后将虚拟头移动到翻转后的第二个节点
以此类推
*/
func swapPairs(head *ListNode) *ListNode {
	h := &ListNode{Next: head}
	t := h
	for t.Next != nil && t.Next.Next != nil {
		a := t.Next
		b := a.Next
		t.Next = b
		a.Next = b.Next
		b.Next = a
		t = a
	}
	return h.Next
}
