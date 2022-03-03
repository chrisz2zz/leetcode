package main

import "fmt"

func main() {
	h := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{
					Val:  4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	r := reverseBetween(h, 2, 4)

	for r != nil {
		fmt.Println(r.Val)
		r = r.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
创建虚拟头节点
定位到要反转的节点和要反转的第一个节点前驱节点(pre)
不断的将第一个节点往后挪
然后把后面的节点放到pre节点的后面
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var pre *ListNode
	h := &ListNode{Next: head}
	count := 0

	t := h
	for t != nil {
		count++
		if count == left {
			pre = t
		}

		if left < count && count <= right {
			tmp := t.Next
			tmp2 := pre.Next
			pre.Next = tmp
			t.Next = tmp.Next
			tmp.Next = tmp2
		} else {
			t = t.Next
		}
	}

	return h.Next
}
