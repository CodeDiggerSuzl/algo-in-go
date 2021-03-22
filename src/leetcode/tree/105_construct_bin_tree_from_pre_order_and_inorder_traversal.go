package tree

/**
根据一棵树的前序遍历与中序遍历构造二叉树。
注意:
你可以假设树中没有重复的元素。

例如，给出
前序遍历 preorder =[3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：
    3
   / \
  9  20
    /  \
   15   7
链接：https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal
*/
func buildTree_1(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	var root int
	for k, v := range inorder {
		if v == preorder[0] {
			root = k
			break
		}
	}
	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree_1(preorder[1:root+1], inorder[:root]),
		Right: buildTree_1(preorder[root+1:], inorder[root+1:]),
	}
}

// https://mmbiz.qpic.cn/sz_mmbiz_jpg/gibkIz0MVqdF8ZItXTVByS26EcqBSS9W6Awr35eI0tibAJ2qW6pDUpgWTv5icgDhRhniaIJg3dpYib7Ph5kqDneL08A/640?wx_fmt=jpeg&tp=webp&wxfrom=5&wx_lazy=1&wx_co=1

/**
func buildTree_1(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree_1(preorder[1:i+1], inorder[:i])
	root.Right = buildTree_1(preorder[i+1:], inorder[i+1:])
	return &root
}

链接：https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/solution/zhi-xing-yong-shi-4-ms-zai-suo-you-go-ti-paax/
*/
