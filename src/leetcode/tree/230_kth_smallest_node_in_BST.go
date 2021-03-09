package tree

/*
给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。
*/
var tHolder []int

// BST 的中序遍历就是排序 递归
func kthSmallest(root *TreeNode, k int) int {
	tHolder = make([]int, 0)
	findKth(root)
	return tHolder[k-1]
}
func findKth(root *TreeNode) {
	if root == nil {
		return
	}
	findKth(root.Left)
	tHolder = append(tHolder, root.Val)
	findKth(root.Right)
}

// golang 为啥连 stack 都没有
