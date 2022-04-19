package main

import (
	"fmt"
)

func main() {
	h1 := &Node{
		Val:  1,
		Next: nil,
	}
	h2 := &Node{
		Val:  3,
		Next: nil,
	}
	h3 := &Node{
		Val:  5,
		Next: nil,
	}

	h1.Next = h2
	h2.Next = h3
	h3.Next = h1

	insert(h2, 0)

	t := h1
	for i := 0; i < 4; i++ {
		fmt.Println(t.Val)
		t = t.Next
	}
}

type Node struct {
	Val  int
	Next *Node
}

func insert(aNode *Node, x int) *Node {
	res := &Node{Val: x}
	if aNode == nil {
		res.Next = res
		return res
	}

	p, en := aNode, aNode
	for {
		if p.Val > p.Next.Val {
			en = p // 此时 en 在最大值的位置
		}

		if p.Val <= x && x <= p.Next.Val {
			p.Next = &Node{x, p.Next}
			break
		}

		// 下面的检查一定要放在上面两个检查之后，不然会漏掉最后一个节点
		if p.Next == aNode {
			en.Next = &Node{x, en.Next}
			break
		}

		p = p.Next
	}

	return aNode
}
