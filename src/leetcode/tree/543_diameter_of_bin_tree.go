package tree

/*

给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。
示例 :
给定二叉树

          1
         / \
        2   3
       / \
      4   5
返回3, 它的长度是路径 [4,2,1,3] 或者[5,2,1,3]。

注意：两结点之间的路径长度是以它们之间边的数目表示。

链接：https://leetcode-cn.com/problems/diameter-of-binary-tree
*/
// 最大左深度和最大右深度的值之和 - 1

func diameterOfBinaryTreeWrong(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ld := getSideDepth(root.Left, 1)
	rd := getSideDepth(root.Right, 1)
	return ld + rd - 1
}

func getSideDepth(node *TreeNode, currDepth int) int {
	if node == nil {
		return currDepth
	}
	currDepth++
	ld := getSideDepth(node.Left, currDepth)
	rd := getSideDepth(node.Right, currDepth)
	if ld >= rd {
		return ld
	}
	return rd

}

// 上面的错了 不是深度，题说了：不一定经过 root 节点，所以不是深度

// 每个节点到不同叶子节的长度之和 -1
// 记录最大的长度 （记录每个节点到叶子的最大值）
var mLen int // 声明

// var mLem = 0 直接这样申请的话不行

func diameterOfBinaryTree(root *TreeNode) int {
	// 赋值（先声明 后赋值）
	mLen = 0
	if root == nil {
		return 0
	}
	getMaxDepth(root)
	return mLen
}

func getMaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ld := getMaxDepth(root.Left)
	rd := getMaxDepth(root.Right)
	if ld+rd >= mLen {
		mLen = ld + rd
	}
	if ld >= rd {
		return ld + 1
	}
	return rd + 1

}
