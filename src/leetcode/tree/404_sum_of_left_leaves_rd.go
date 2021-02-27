package tree

/*
计算给定二叉树的所有左叶子之和。

示例：

    3
   / \
  9  20
    /  \
   15   7

在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24

链接：https://leetcode-cn.com/problems/sum-of-left-leaves
*/

// 审题失误 这个方法是求得左节点之和，但是要的是左叶子之和
var leftSum = 0

func sumOfLeftNode(root *TreeNode) int {
    if root == nil {
        return 0
    }
    if root.Left != nil {
        leftSum += root.Left.Val
    }
    sumOfLeftNode(root.Left)
    sumOfLeftNode(root.Right)
    return leftSum
}

// 求左叶子🍃节点之和
func sumOfLeftLeaves(root *TreeNode) (r int) {
    // if root == nil {
    //     return 0
    // }
    sumLevees(root, &r)
    return r
}

func sumLevees(root *TreeNode, r *int) {
    if root == nil {
        return
    }
    if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
        *r += root.Left.Val
    }
    sumLevees(root.Left, r)
    sumLevees(root.Right, r)
}
