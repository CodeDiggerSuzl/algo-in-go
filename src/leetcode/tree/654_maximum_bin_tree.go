package tree

import "math"

/*
给定一个不含重复元素的整数数组 nums 。一个以此数组直接递归构建的 最大二叉树 定义如下：

二叉树的根是数组 nums 中的最大元素。
左子树是通过数组中 最大值左边部分 递归构造出的最大二叉树。
右子树是通过数组中 最大值右边部分 递归构造出的最大二叉树。
返回有给定数组 nums 构建的 最大二叉树 。

链接：https://leetcode-cn.com/problems/maximum-binary-tree
*/
func constructMaximumBinaryTree(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	low := -1
	// 找到 root 的节点 前序遍历(只能）
	min := math.MinInt64
	for idx, v := range arr {
		if v > min {
			low = idx
			min = v
		}
	}
	// find root
	node := &TreeNode{min, nil, nil}
	node.Left = constructMaximumBinaryTree(arr[:low])    // root.left 递归
	node.Right = constructMaximumBinaryTree(arr[low+1:]) // root.right
	return node
}
