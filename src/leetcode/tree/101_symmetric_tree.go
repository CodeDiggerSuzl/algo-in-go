package tree

/*

同 剑指 Offer 28. 对称的二叉树: https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof/

给定一个二叉树，检查它是否是镜像对称的。
例如，二叉树[1,2,2,3,4,4,3] 是对称的。
    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个[1,2,2,null,3,null,3] 则不是镜像对称的:
    1
   / \
  2   2
   \   \
   3    3
链接：https://leetcode-cn.com/problems/symmetric-tree
*/
// 1. 迭代法
// left.val == r.val
// 左树的左节点的值 == 右树的右节点的值
// 左树的右节点的值 == 右树的左节点的值
func isSymmetric(root *TreeNode) bool {
	return checkWithTwoNodes(root, root)
}

func checkWithTwoNodes(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	// if l.Val != r.Val {
	// return false
	// }
	return l.Val == r.Val && checkWithTwoNodes(l.Left, r.Right) && checkWithTwoNodes(l.Right, r.Left)
}

// 迭代法 层序遍历 每一层级的 node 数量应为偶数,且对称
