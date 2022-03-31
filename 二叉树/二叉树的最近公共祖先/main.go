/**
* @Author: Chao
* @Date: 2022/3/31 12:56
* @Version: 1.0
 */

package main

import "fmt"

func main() {
	tree := &TreeNode{
		Val:   3,
		Left:  &TreeNode{
			Val:   5,
			Left:  &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   2,
				Left:  &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   8,
				Left:  nil,
				Right: nil,
			},
		},
	}

	fmt.Println(lowestCommonAncestor(tree, &TreeNode{Val: 7}, &TreeNode{Val: 4}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//p 和 q 均存在于给定的二叉树中
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	//前序遍历, 根左右
	//找到其中一个节点返回
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	//在该节点的左右子树找到了p,q,找到则为公共节点
	if left != nil && right != nil {
		return root
	}

	//当前节点的左子树中只找到了p或者q
	//说明没找到的p或者q一定在改节点的左右子树中
	//所以当前节点既是p或者q,又是p和q的公共祖先
	if left != nil {
		return left
	}

	//同left
	return right
}
