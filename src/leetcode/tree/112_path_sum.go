package tree

/*
 给你二叉树的根节点root 和一个表示目标和的整数targetSum ，

判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和targetSum 。

叶子节点 是指没有子节点的节点。

链接：https://leetcode-cn.com/problems/path-sum
*/
// 一条路径上的值 明日次都减去上个 node 的 Val （🤒）
func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }
    if root.Left == nil && root.Right == nil {
        return targetSum == root.Val
    }
    diff := targetSum - root.Val
    // 每一条路径 要么是左 要么是右
    return hasPathSum(root.Left, diff) || hasPathSum(root.Right, diff)
}
