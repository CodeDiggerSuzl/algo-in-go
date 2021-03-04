package tree

import "math"

/*
您需要在二叉树的每一行中找到最大的值。
示例：
输入:
          1
         / \
        3   2
       / \   \
      5   3   9
输出: [1, 3, 9]

链接：https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row
*/
var r []int

func largestValues(root *TreeNode) []int {
	r = make([]int, 0)
	findLargest(root, 0)
	return r
}
func findLargest(n *TreeNode, level int) {
	if n == nil {
		return
	}

	// !核心在这四行代码
	if len(r) == level {
		r = append(r, math.MinInt64)
	}
	if n.Val > r[level] {
		r[level] = n.Val
	}

	findLargest(n.Left, level+1)
	findLargest(n.Right, level+1)
}
