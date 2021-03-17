package tree

/*
https://leetcode-cn.com/problems/binary-tree-right-side-view/

给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
*/
// 层序遍历返回 arr 最后的节点
var rightView []int

func rightSideView(root *TreeNode) []int {
	rightView = make([]int, 0)
	getRightView(root, 0)
	return rightView
}
func getRightView(n *TreeNode, lv int) {
	if n == nil {
		return
	}

	if lv == len(rightView) {
		rightView = append(rightView, n.Val)
	}
	getRightView(n.Right, lv+1)
	getRightView(n.Left, lv+1)

}

// 思路： 我们按照 「根结点 -> 右子树 -> 左子树」 的顺序访问，就可以保证每层都是最先访问最右边的节点的。
//（与先序遍历 「根结点 -> 左子树 -> 右子树」 正好相反，先序遍历每层最先访问的是最左边的节点）

// 使用数组指针
func rightSideViewPtr(root *TreeNode) (r []int) {
	getRightViewWithPtr(root, 0, &r)
	return // r
}
func getRightViewWithPtr(n *TreeNode, lv int, r *[]int) {
	if n == nil {
		return
	}

	if lv == len(*r) {
		*r = append(*r, n.Val)
	}
	getRightViewWithPtr(n.Right, lv+1, r)
	getRightViewWithPtr(n.Left, lv+1, r)
}
