package bintree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
// 本质是前序遍历
var res [][]int

func levelOrder(root *TreeNode) [][]int {
	res = make([][]int, 0)

	levelTraversal(root, 0)
	return res
}

func levelTraversal(root *TreeNode, level int) {
	if root == nil {
		return
	}
	// 每一层加一个数组
	if level == len(res) {
		res = append(res, []int{})
	}
	res[level] = append(res[level], root.Val)
	levelTraversal(root.Left, level+1)
	levelTraversal(root.Right, level+1)

}
