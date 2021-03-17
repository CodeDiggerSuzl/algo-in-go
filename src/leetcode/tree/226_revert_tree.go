package tree

/*
翻转一棵二叉树。
示例：
输入：

     4
   /   \
  2     7
 / \   / \
1   3 6   9
输出：
     4
   /   \
  7     2
 / \   / \
9   6 3   1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/invert-binary-tree
*/
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
