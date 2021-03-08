package tree

/*
给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。

完全二叉树 的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~ 2h 个节点。

链接：https://leetcode-cn.com/problems/count-complete-tree-nodes
*/

// 第一种方法;效率较低,但是比较通用
var cntMap map[int]int

func countNodes(root *TreeNode) int {
	cntMap = make(map[int]int, 0)
	if root == nil {
		return 0
	}
	levelOrderCnt(root, 0)
	cnt := 0
	for _, v := range cntMap {
		cnt += v
	}
	return cnt
}
func levelOrderCnt(node *TreeNode, level int) {
	if node == nil {
		return
	}
	if _, ok := cntMap[level]; ok {
		cntMap[level] ++
	} else {
		cntMap[level] = 1
	}
	levelOrderCnt(node.Left, level+1)
	levelOrderCnt(node.Right, level+1)
}

// 递归解法 better
func countNodes2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := countNodes2(root.Left)
	r := countNodes2(root.Right)
	return r + l + 1
}
