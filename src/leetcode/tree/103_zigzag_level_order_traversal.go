package tree

/**
给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回锯齿形层序遍历如下：

[
  [3],
  [20,9],
  [15,7]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal
*/
var zigRes [][]int

func zigzagLevelOrder(root *TreeNode) [][]int {
	zigRes = make([][]int, 0)
	zigDfs(root, 0)
	for i, v := range zigRes {
		if i%2 != 0 {
			reverseSlice(v)
		}
	}
	return zigRes
}

func reverseSlice(v []int) {
	for i, j := 0, len(v)-1; i < j; i, j = i+1, j-1 {
		v[i], v[j] = v[j], v[i]
	}
}
func zigDfs(n *TreeNode, lv int) {
	if n == nil {
		return
	}
	if len(zigRes) == lv {
		zigRes = append(zigRes, []int{})
	}
	zigRes[lv] = append(zigRes[lv], n.Val)
	zigDfs(n.Left, lv+1)
	zigDfs(n.Right, lv+1)
}
