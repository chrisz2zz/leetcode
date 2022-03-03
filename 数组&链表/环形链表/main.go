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

	fmt.Println(hasCycle(l1))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
快慢指针
声明两个指针,一个移动一步,一个移动两步
若相遇得有环
若slow或者fast或者fast.next为nil
说明无环
 */
func hasCycle2(head *ListNode) bool {
	slow, fast := head, head
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

/*
记录法
循环遍历,把每个节点地址都记录下来保存到map中
只要遇到map中存在即有环
 */
func hasCycle(head *ListNode) bool {
	index := make(map[*ListNode]bool)
	for head != nil {
		if index[head] {
			return true
		}
		index[head] = true
		head = head.Next
	}
	return false
}