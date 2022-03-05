package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	prev := &ListNode{
		Val:  0,
		Next: head,
	}

	t := prev

	for head != nil {
		if head.Val == val {
			t.Next = head.Next
		}

		head = head.Next
		t = t.Next
	}
	return prev.Next
}
