package tree

import "math"

// rd
// 高度和深度的定义
/*
给定一个二叉树，判断它是否是高度平衡的二叉树。
本题中，一棵高度平衡二叉树定义为：
一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。
https://leetcode-cn.com/problems/balanced-binary-tree/
*/

func isBalanced(root *TreeNode) bool {
	if getBlanceDepth(root) == -1 {
		return false
	}
	return true
}

// 要后续,最后用根的值
// -1 表示不是平衡二叉树了
func getBlanceDepth(n *TreeNode) int {
	// 终止条件 返回 0 表示当前节点为根节点树的高度为 0
	if n == nil {
		return 0
	}
	ld := getBlanceDepth(n.Left) // 左
	if ld == -1 {
		return -1
	}
	rd := getBlanceDepth(n.Right) // 右
	if rd == -1 {
		return -1
	}
	if math.Abs(float64(ld-rd)) > 1 {
		return -1
	}
	if ld > rd {
		return ld + 1
	}
	return rd + 1
}

// 节点的高度和深度的区别
