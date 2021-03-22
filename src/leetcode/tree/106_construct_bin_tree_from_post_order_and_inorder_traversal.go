package tree

/**
根据一棵树的前序遍历与中序遍历构造二叉树。
注意:
你可以假设树中没有重复的元素。

例如，给出
前序遍历 preorder =[3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：
    3
   / \
  9  20
    /  \
   15   7
链接：https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal
*/
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	var idx int
	for k, v := range inorder {
		if v == postorder[len(postorder)-1] {
			idx = k
			break
		}
	}
	return &TreeNode{
		Val:   postorder[len(postorder)-1],
		Left:  buildTree(inorder[:idx], postorder[:idx]),
		Right: buildTree(inorder[idx+1:], postorder[idx:len(postorder)-1]),
	}
}

// !! 根据前中后 写出 node 的顺序 ！！
