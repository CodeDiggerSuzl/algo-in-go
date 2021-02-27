package tree

/*
给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。
你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，否则不为NULL 的节点将直接作为新二叉树的节点。
示例1:
输入:
	Tree 1                     Tree 2
          1                         2
         / \                       / \
        3   2                     1   3
       /                           \   \
      5                             4   7
输出:
合并后的树:
	     3
	    / \
	   4   5
	  / \   \
	 5   4   7
注意:合并必须从两个树的根节点开始。

链接：https://leetcode-cn.com/problems/merge-two-binary-trees
*/
// 往一个🌲上和并
func mergeTrees(r1, r2 *TreeNode) *TreeNode {

    if r1 == nil {
        return r2
    }
    if r2 == nil {
        return r1
    }
    // r1.Val += r2.Val 三个位置都可以
    r1.Left = mergeTrees(r1.Left, r2.Left)
    r1.Val += r2.Val
    r1.Right = mergeTrees(r1.Right, r2.Right)

    // r1.Val += r2.Val
    return r1
}
