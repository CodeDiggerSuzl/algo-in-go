package tree

/*
给定一个 N 叉树，返回其节点值的层序遍历。（即从左到右，逐层遍历）。
树的序列化输入是用层序遍历，每组子节点都由 null 值分隔（参见示例）。

https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/
*/
var nAryLevelRes [][]int

func nAryLevelOrder(root *NNode) [][]int {
	nAryLevelRes = make([][]int, 0)
	nOrderTraversal(root, 0)
	return nAryLevelRes
}

func nOrderTraversal(node *NNode, level int) {
	if node == nil {
		return
	}
	if level == len(nAryLevelRes) {
		nAryLevelRes = append(nAryLevelRes, []int{})
	}
	nAryLevelRes[level] = append(nAryLevelRes[level], node.Val)
	for _, child := range node.Children {
		nOrderTraversal(child, level+1)
	}
}
