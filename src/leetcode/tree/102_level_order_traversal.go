package tree

/*
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
示例：
二叉树：[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层序遍历结果：

[
  [3],
  [9,20],
  [15,7]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal
*/

// 本质是前序遍历
var res [][]int

// levelOrder
func levelOrder(root *TreeNode) [][]int {
	res = make([][]int, 0)
	levelTraversal(root, 0)
	return res
}

func levelTraversal(root *TreeNode, level int) {
	if root == nil {
		return
	}
	// 每一层加一个数组
	if level == len(res) {
		res = append(res, []int{})
	}
	res[level] = append(res[level], root.Val)
	levelTraversal(root.Left, level+1)
	levelTraversal(root.Right, level+1)

}

// 层序遍历的标准
