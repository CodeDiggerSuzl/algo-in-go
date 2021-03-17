package tree

import "math"

/**
给定一个非空特殊的二叉树，每个节点都是正数，并且每个节点的子节点数量只能为2或0。如果一个节点有两个子节点的话，那么该节点的值等于两个子节点中较小的一个。

更正式地说，root.val = min(root.left.val, root.right.val) 总成立。
给出这样的一个二叉树，你需要输出所有节点中的第二小的值。如果第二小的值不存在的话，输出 -1 。
链接：https://leetcode-cn.com/problems/second-minimum-node-in-a-binary-tree
*/

// 最小值是 root
var secMin int
var rootVal int

func findSecondMinimumValue(root *TreeNode) int {
	secMin = math.MaxInt64
	if root == nil {
		return -1
	}
	rootVal = root.Val
	// 最小值是 root 节点
	dfs(root)
	// 没有发生变化
	if secMin == math.MaxInt64 {
		return -1
	}
	return secMin
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Val != rootVal && root.Val <= secMin {
		secMin = root.Val
	}
	dfs(root.Left)
	dfs(root.Right)
}
