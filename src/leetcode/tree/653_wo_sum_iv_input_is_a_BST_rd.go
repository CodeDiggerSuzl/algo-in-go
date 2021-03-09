package tree

/*
给定一个二叉搜索树和一个目标结果，如果 BST 中存在两个元素且它们的和等于给定的目标结果，则返回 true。
https://leetcode-cn.com/problems/two-sum-iv-input-is-a-bst/
*/
// 基本解法就是 遍历加第一题的 tow sum

// 进阶解法(但是感觉没用到 BST 的特性,而且使用了额外的数据结构
func findTarget(root *TreeNode, k int) bool {
	tarMap := map[int]bool{}
	return findTargetDfs(root,k,tarMap)
}
func findTargetDfs(root *TreeNode, k int, tarMap map[int]bool) bool {
	if root == nil {
		return false
	}
	if _,ok := tarMap[k-root.Val];ok{
		return true
	}
	tarMap[root.Val] = true
	return findTargetDfs(root.Left,k,tarMap)||findTargetDfs(root.Right,k,tarMap)
}

