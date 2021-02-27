package tree

import "ds-algo/src/utils"

/*

输入一棵二叉树的根节点，求该树的深度。从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。

例如：

给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。

https://leetcode-cn.com/problems/er-cha-shu-de-shen-du-lcof/
*/
/*

class Solution {
public:
    int maxDepth(TreeNode* root) {
        return root != nullptr ? (max(maxDepth(root->left), maxDepth(root->right)) + 1): 0;
    }
};
*/
// 
// ! 前序遍历  左子树或者右子树的最大值 better
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    lmd := maxDepth(root.Left)
    rmd := maxDepth(root.Right)
    return utils.MaxInt(lmd, rmd) + 1 // 加上 root 那一层
}

// 方法二 就的
func getMaxDepthWithRecode(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return getDepth(root, 0)

}
func getDepth(node *TreeNode, depth int) int {
    if node == nil {
        return depth
    }
    depth++
    return utils.MaxInt(getDepth(node.Left, depth), getDepth(node.Right, depth))
}
