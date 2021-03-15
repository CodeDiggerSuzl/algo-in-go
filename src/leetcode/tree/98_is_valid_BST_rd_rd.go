package tree

/*
给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/validate-binary-search-tree
*/
// 初次尝试 ❌ 解法: 我们要⽐较的是 左⼦树所有节点⼩于中间节点，右⼦树所有节点⼤于中间节点。所以以上代码的判断逻辑是错误
//func isValidBST3(root *TreeNode) bool {
//	if root == nil {
//		return true
//	}
//	if root.Left == nil {
//		return isValidBST3(root.Right)
//	}
//	if root.Right == nil {
//		return isValidBST3(root.Left)
//	}
//	if root.Left.Val >= root.Val || root.Right.Val <= root.Val {
//		return false
//	}
//	return isValidBST3(root.Left) || isValidBST3(root.Right)
//}

// ------------------------------------------------------
// 利用 中序遍历有序 也不对 shit
var carry *TreeNode
func isValidBST3(root *TreeNode) bool {
	carry = &TreeNode{} // fuck  ..........................
	return check(root)
}

// still not right, shit
func check(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var l = check(root.Left)
	// 中序遍历
	if carry != nil && carry.Val >= root.Val {
		return false
	}
	carry = root
	var  r = check(root)
	return l && r
}

/* java 这样都可以 所以上面的是因为 nil 的问题
class Solution {
    TreeNode p = null;
    public boolean isValidBST(TreeNode root) {
        if (root == null){
            return true;
        }
        boolean l = isValidBST(root.left);
        if (p!=null && p.val >= root.val){
            return false;
        }
        p = root;
        boolean r = isValidBST(root.right);
        return l && r;
    }
}
*/

// BST 的所有左子树的值都小于 root,右子树的值都大于 root, 这种将约束传递给子树的所有节点也是一种二叉树的解题技巧
func isValidBST(root *TreeNode) bool {
	return isValidBSTCarryWithValue(root, nil, nil)
}

func isValidBSTCarryWithValue(root *TreeNode, min *TreeNode, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}
	return isValidBSTCarryWithValue(root.Left, min, root) && isValidBSTCarryWithValue(root.Right, root, max)
}
