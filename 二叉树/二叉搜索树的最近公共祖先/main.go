/**
* @Author: Chao
* @Date: 2022/3/31 13:16
* @Version: 1.0
 */

package main

import "fmt"

func main() {
	ast := &TreeNode{
		Val:   6,
		Left:  &TreeNode{
			Val:   2,
			Left:  &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val:   8,
			Left:  &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
		},
	}

	fmt.Println(lowestCommonAncestor(ast,
		&TreeNode{Val: 2}, &TreeNode{Val: 8}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//所有节点的值都是唯一的。
//p、q 为不同节点且均存在于给定的二叉搜索树中。
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	max, min := finMaxMin(p, q)

	return find(root, max, min)
}

func find(root *TreeNode, max, min int) *TreeNode {
	//叶子结点处理
	if root == nil {
		return nil
	}

	//前序遍历, 根左右
	//因为p q都在树中
	//所以只需要找到一个节点
	//大于等于小的那个节点且小于等于大的那个节点
	//即为公共祖先
	if min <= root.Val && root.Val <= max {
		return root
	}

	//遍历左子树
	left := find(root.Left, max, min)
	if left != nil {
		return left
	}
	//遍历右子树
	right := find(root.Right, max, min)

	return right
}

func finMaxMin(p, q *TreeNode) (max, min int) {
	if p.Val > q.Val {
		return p.Val, q.Val
	}

	return q.Val, p.Val
}