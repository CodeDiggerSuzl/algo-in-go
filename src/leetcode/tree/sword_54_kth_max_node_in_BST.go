package tree

/*
给定一棵二叉搜索树，请找出其中第k大的节点。
https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/
*/
var kMax []int

// BST 中序就是 排序
func kthLargest(root *TreeNode, k int) int {
	kMax := make([]int, 0)
	findKthMax(root)
	return kMax[len(kMax)-k]
}

func findKthMax(root *TreeNode) {
	if root == nil {
		return
	}
	findKthMax(root.Left)
	kMax = append(kMax, root.Val)
	findKthMax(root.Right)
}
