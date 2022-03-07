package main

import "fmt"

func main() {
	l1 := &ListNode{
		Val:  3,
		Next: nil,
	}
	l2 := &ListNode{
		Val:  2,
		Next: nil,
	}
	l3 := &ListNode{
		Val:  0,
		Next: nil,
	}
	l4 := &ListNode{
		Val:  -4,
		Next: nil,
	}

	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l2

	fmt.Println(detectCycle(l1))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
map记录法
*/
func detectCycle2(head *ListNode) *ListNode {
	index := make(map[*ListNode]bool)
	for head != nil {
		if index[head] {
			return head
		}
		index[head] = true
		head = head.Next
	}
	return nil
}

/*
快慢指针
*/
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow := head
	fast := head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}

	return nil
}
