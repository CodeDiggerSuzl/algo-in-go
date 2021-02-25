package bintree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// https://leetcode-cn.com/problems/invert-binary-tree/ 翻转二叉树
// 本质是二叉树的前序或者后续遍历(不能是中序)
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	var tmp *TreeNode
	tmp = root.Left
	root.Left = root.Right
	root.Right = tmp

	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
