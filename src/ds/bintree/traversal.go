package bintree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历 root -> left -> right
/*
	递归式前序遍历: root -> left -> right
*/
func preorderTraversal(root *TreeNode) []int {
	var result []int
	var preOrder func(*TreeNode)

	preOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, root.Val)
		preOrder(root.Left)
		preOrder(root.Right)
	}
	preOrder(root)
	return result
}

// 中序遍历 左-> 根-> 右
func inorderTraversal(root *TreeNode) []int {
	var res []int
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return res
}

// 后序遍历 左 -> 右 -> 根
func postorderTraversal(root *TreeNode) []int {
	var res []int
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		inorder(node.Right)
		res = append(res, node.Val)
	}
	inorder(root)
	return res
}
