package tree

// https://leetcode-cn.com/problems/same-tree/
// 给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。
//
// 如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

// 本质是前序遍历
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }
    if p == nil || q == nil {
        return false
    }
    if q.Val != p.Val {
        return false
    }
    return isSameTree(p.Right, q.Right) && isSameTree(p.Left, q.Left)
}