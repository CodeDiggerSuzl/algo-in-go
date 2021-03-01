package tree

/**
给定一个 N 叉树，找到其最大深度。

最大深度是指从根节点到最远叶子节点的最长路径上的节点总数。

N 叉树输入按层序遍历序列化表示，每组子节点由空值分隔（请参见示例）。

链接：https://leetcode-cn.com/problems/maximum-depth-of-n-ary-tree
*/

// maxNryTreeDepth 类似二叉树深度
func maxNryTreeDepth(root *NNode) int {
	if root == nil {
		return 0
	}
	eachMaxDepth := 0 // 找到 n 个子节点中的最大深度
	for _, n := range root.Children {
		max := maxNryTreeDepth(n)
		if max > eachMaxDepth {
			eachMaxDepth = max
		}
	}
	return eachMaxDepth + 1
}
