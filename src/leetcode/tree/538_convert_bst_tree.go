package tree

/**
给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），使每个节点 node 的新值等于原树中大于或等于 node.val 的值之和。
https://leetcode-cn.com/problems/convert-bst-to-greater-tree/
*/
var sum int

func convertBST(root *TreeNode) *TreeNode {
	sum = 0
	f(root)
	return root
}

// 先右后左
func f(r *TreeNode) {
	if r == nil {
		return
	}
	f(r.Right)
	sum = sum + r.Val
	r.Val = sum
	f(r.Left)
}
