package tree

/**
给定一个非空二叉树, 返回一个由每层节点平均值组成的数组。
示例 1：
输入：
    3
   / \
  9  20
    /  \
   15   7
输出：[3, 14.5, 11]
解释：
第 0 层的平均值是 3 ,  第1层是 14.5 , 第2层是 11 。因此返回 [3, 14.5, 11] 。

提示：
	节点值的范围在32位有符号整数范围内。

链接：https://leetcode-cn.com/problems/average-of-levels-in-binary-tree
*/

// 中序遍历,最后求平均值
var averageOfLevelsTemp [][]int

func averageOfLevels(root *TreeNode) (r []float64) {
	averageOfLevelsTemp = make([][]int, 0)
	averageLevel(root, 0)
	result := make([]float64, 0)
	// 求和 然后求平均值
	for _, levelArray := range averageOfLevelsTemp {
		eachLevelSum := 0
		for _, i := range levelArray {
			eachLevelSum += i
		}
		result = append(result, float64(eachLevelSum)/float64(len(levelArray)))
	}
	return result
}

// 中序遍历
func averageLevel(node *TreeNode, level int) {
	if node == nil {
		return
	}
	if level == len(averageOfLevelsTemp) {
		averageOfLevelsTemp = append(averageOfLevelsTemp, []int{})
	}
	averageOfLevelsTemp[level] = append(averageOfLevelsTemp[level], node.Val)
	averageLevel(node.Left, level+1)
	averageLevel(node.Right, level+1)
}
