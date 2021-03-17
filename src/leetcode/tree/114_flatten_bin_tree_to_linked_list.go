package tree

/*
* good one !
给你二叉树的根结点 root ，请你将它展开为一个单链表：

展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。

链接：https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list
*/
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	// 后续遍历
	flatten(root.Left)
	flatten(root.Right)
	// 到这里已经拉平了左树和右树
	temp := root.Right     // 暂存右树
	root.Right = root.Left // 右树指向左树
	root.Left = nil
	for root.Right != nil {
		root = root.Right // 找到右边最后边的 node
	}
	root.Right = temp // 接上原来的右边节点
}
