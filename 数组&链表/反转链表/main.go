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
//	rs := reverseList(head)
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
////递归版本
////func reverseList(head *ListNode) *ListNode {
////	//判断结束条件，head为nil或者head.Next为nil直接返回
////	if head == nil || head.Next == nil {
////		return head
////	}
////
////	//递归进入下一个节点
////	next := reverseList(head.Next)
////
////	//将当前节点的子节点的Next指向自己,相当于4->5->nil变成4->5->4
////	head.Next.Next = head
////	//删除当前节点的子节点,相当于4->5->4 变成 5->4->nil
////	head.Next = nil
////
////	//返回最后节点
////	return next
////}
//
////迭代版本
//func reverseList(head *ListNode) *ListNode {
//	//记录前驱节点
//	var pre *ListNode
//
//	//迭代链表
//	for head != nil {
//		//记录当前节点的子节点
//		next := head.Next
//		//当前节点指向前驱节点
//		head.Next = pre
//		//前驱节点后移
//		pre = head
//		//当前节点后移
//		head = next
//	}
//
//	return pre
//}

func main() {
	h := &ListNode{
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

	r := reverseList(h)

	for r != nil {
		fmt.Println(r.Val)
		r = r.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList2(head *ListNode) *ListNode {
	var cur, prev *ListNode = head, nil
	for cur != nil {
		cur.Next, prev, cur = prev, cur, cur.Next
	}
	return prev
}

/*
类似于双游标
第一个游标为nil  表示head之前的元素
第二个游标为head
然后取出第二个游标.next保存
再将第二个游标.next指向第一个游标
再将第一个游标向后移,第二游标向后移
直到第二个游标为nil,此时第一个游标为原链表最后一个节点
*/
//func reverseList(head *ListNode) *ListNode {
//	var cur, prev *ListNode = head, nil
//	for cur != nil {
//		tmp := cur.Next
//		cur.Next = prev
//		prev = cur
//		cur = tmp
//	}
//	return prev
//}

// 游标法 review
func reverseList3(head *ListNode) *ListNode {
	var cur, prev *ListNode
	cur = head
	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}
	return prev
}

//穿针法
/*
	制造一个虚拟头结点(prev),prev的next指向head
	然后取出head的next并记录(tmp)
	将head.Next指向tmp的Next
	将tmp.Next指向prev.next
	再将prev.Next指向tmp

	直到head.Next为nil 或 head为nil

	最后返回prev.Next
*/
func reverseList(head *ListNode) *ListNode {
	prev := ListNode{}

	prev.Next = head

	for head != nil && head.Next != nil {
		tmp := head.Next
		head.Next = tmp.Next
		tmp.Next = prev.Next
		prev.Next = tmp
	}

	return prev.Next
}
