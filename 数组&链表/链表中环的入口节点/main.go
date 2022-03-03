/**
* @Author: Chao
* @Date: 2022/3/3 12:57
* @Version: 1.0
 */

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

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			fast = head
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return fast
		}
	}

	return nil
}
