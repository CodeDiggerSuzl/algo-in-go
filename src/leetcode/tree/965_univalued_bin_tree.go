package tree

/**
如果二叉树每个节点都具有相同的值，那么该二叉树就是单值二叉树。

只有给定的树是单值二叉树时，才返回true；否则返回 false。
示例 1：
输入：[1,1,1,1,1,null,1]
输出：true

示例 2：
输入：[2,2,2,5,2]
输出：false
链接：https://leetcode-cn.com/problems/univalued-binary-tree
*/
// 理解递归: 都是同值： 左树是同值 && 右树是同值

// isUnivalTree judge the binary tree has the same value or not
func isUnivalTree(root *TreeNode) bool {
    // good 为啥想不到  （判断条件 root.val == root.next.val）
    lb := root.Left == nil || (root.Val == root.Left.Val && isUnivalTree(root.Left))
    rb := root.Right == nil || (root.Val == root.Right.Val && isUnivalTree(root.Right))
    return lb && rb
}
