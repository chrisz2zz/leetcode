package main

import "fmt"

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	r := removeNthFromEnd(head, 2)

	for r != nil {
		fmt.Println(r.Val)
		r = r.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 借用数组
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	hair := &ListNode{Next: head}

	list := make([]*ListNode, 0)
	head = hair
	for head != nil {
		list = append(list, head)
		head = head.Next
	}

	if len(list) < n {
		return nil
	}

	index := len(list) - n

	node := list[index]

	prev := list[index-1]

	prev.Next = node.Next

	node = nil

	return hair.Next
}

//快慢指针
func removeNthFromEnd3(head *ListNode, n int) *ListNode {
	hair := &ListNode{Next: head}
	slow, fast, prev := hair, hair, hair

	count := 0
	for slow != nil {
		if count >= n {
			prev = fast
			fast = fast.Next
		}

		slow = slow.Next
		count++
	}

	prev.Next = fast.Next
	fast = nil
	return hair.Next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	prev := &ListNode{Next: head}
	fast, slow := prev, prev

	var tmp *ListNode
	counter := 0
	for fast.Next != nil {
		fast = fast.Next
		counter++
		if counter >= n {
			tmp = slow
			slow = slow.Next
		}
	}

	tmp.Next = slow.Next

	return prev.Next
}