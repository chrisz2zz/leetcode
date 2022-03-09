package main

import "fmt"

func main() {
	h1 := &Node{
		Val:    7,
		Next:   nil,
		Random: nil,
	}

	h2 := &Node{
		Val:    13,
		Next:   nil,
		Random: nil,
	}

	h3 := &Node{
		Val:    11,
		Next:   nil,
		Random: nil,
	}

	h4 := &Node{
		Val:    10,
		Next:   nil,
		Random: nil,
	}

	h5 := &Node{
		Val:    1,
		Next:   nil,
		Random: nil,
	}

	h1.Next = h2
	h2.Next = h3
	h3.Next = h4
	h4.Next = h5
	h2.Random = h1
	h3.Random = h5
	h4.Random = h3
	h5.Random = h1

	c := copyRandomList(h1)

	for c != nil {
		fmt.Println(c.Val, c.Random)
		c = c.Next
	}
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	hair := &Node{}

	p := head
	//t := hair
	for p != nil {

	}

	return hair.Next
}
