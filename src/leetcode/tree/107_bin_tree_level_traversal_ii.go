package tree

/**
给定一个二叉树，返回其节点值自底向上的层序遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

例如：
给定二叉树 [3,9,20,null,null,15,7],
返回其自底向上的层序遍历为：

[
  [15,7],
  [9,20],
  [3]
]

链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii
*/
var bottomRes [][]int

func levelOrderBottom2(root *TreeNode) [][]int {
	bottomRes = make([][]int, 0)
	levelTraversalBottom(root, 0)
	length := len(bottomRes) // 3
	final := make([][]int, length)
	for i, item := range bottomRes {
		// 3
		// 0 2
		// 1 1
		// 2 0
		final[length-i-1] = item
	}
	return final
}

func levelTraversalBottom(node *TreeNode, l int) {
	if node == nil {
		return
	}
	if len(bottomRes) == l {
		bottomRes = append(bottomRes, []int{})
	}
	bottomRes[l] = append(bottomRes[l], node.Val)
	levelTraversalBottom(node.Left, l+1)
	levelTraversalBottom(node.Right, l+1)
}
