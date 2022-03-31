/**
* @Author: Chao
* @Date: 2022/3/31 13:31
* @Version: 1.0
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	ast1 := &TreeNode{
		Val:   2,
		Left:  &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	ast2 := &TreeNode{
		Val:   5,
		Left:  &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 6},
		},
	}

	ast3 := &TreeNode{
		Val:   0,
		Left:  &TreeNode{
			Val:   -1,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}
	fmt.Println(isValidBST(ast3))
	fmt.Println(isValidBST(ast1), isValidBST(ast2))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return judge(root, math.MinInt, math.MaxInt)
}

func judge(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}

	if root.Val <= lower || root.Val >= upper {
		return false
	}

	return judge(root.Left, lower, root.Val) && judge(root.Right, root.Val, upper)
}