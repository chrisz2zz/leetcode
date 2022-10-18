/**
* @Author: Chao
* @Date: 2022/8/17 21:18
* @Version: 1.0
 */

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	head := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val:  6,
				Left: nil,
				Right: &TreeNode{
					Val:   8,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}
	fmt.Println(deepestLeavesSum(head))
}

func deepestLeavesSum(root *TreeNode) int {
	record := make(map[int]int)

	traverse(root, 0, record)

	max := 0
	for k := range record {
		if k > max {
			max = k
		}
	}

	return record[max]
}

func traverse(node *TreeNode, depth int, record map[int]int) {
	if node == nil {
		return
	}

	record[depth] += node.Val

	traverse(node.Left, depth+1, record)
	traverse(node.Right, depth+1, record)
}
