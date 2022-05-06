package main

import "fmt"

/*
	所为广度优先
	就是每次把一层的节点放入一个队列
	然后遍历
	再放入再遍历
	直到找到满足条件
*/

func main() {
	tree := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}

	fmt.Println(minDepth(&tree))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	ts := make([]*TreeNode, 0)
	depths := make([]int, 0)

	if root != nil {
		ts = append(ts, root)
		depths = append(depths, 1)
	}

	for i := 0; i < len(ts); i++ {
		if ts[i].Left == nil && ts[i].Right == nil {
			return depths[i]
		}

		if ts[i].Left != nil {
			ts = append(ts, ts[i].Left)
			depths = append(depths, depths[i]+1)
		}

		if ts[i].Right != nil {
			ts = append(ts, ts[i].Right)
			depths = append(depths, depths[i]+1)
		}
	}

	return 0
}
