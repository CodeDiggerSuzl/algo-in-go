package tree

import "strconv"

var memo map[string]int
var result652 []*TreeNode

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	memo = make(map[string]int, 0)
	result652 = make([]*TreeNode, 0)
	findDuplicateSubTree(root)
	return result652
}

func findDuplicateSubTree(root *TreeNode) string {
	if root == nil {
		return "*"
	}
	left := findDuplicateSubTree(root.Left)
	right := findDuplicateSubTree(root.Right)
	str := left + "-" + right + "-" + strconv.Itoa(root.Val)
	var freq int
	if v, ok := memo[str]; ok {
		freq = v
		if v == 1 {
			result652 = append(result652, root)
		}
	}
	memo[str] = freq + 1
	return str
}
