package tree

// https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/
// 剑指 Offer 27. 二叉树的镜像
/*
请完成一个函数，输入一个二叉树，该函数输出它的镜像。
例如输入：

     4
   /   \
  2     7
 / \   / \
1   3 6   9
镜像输出：
     4
   /   \
  7     2
 / \   / \
9   6 3   1
*/

// 本质也是前序、后序遍历
func mirrorTree(root *TreeNode) *TreeNode {
    if root == nil {
        return root
    }
    var temp *TreeNode
    temp = root.Left
    root.Left = root.Right
    root.Right = temp
    mirrorTree(root.Left)
    mirrorTree(root.Right)
    return root
}
