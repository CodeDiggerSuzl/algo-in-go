package tree

// 给定一个 N 叉树，返回其节点值的前序遍历。
// https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/

var resultArr []int

func preorder(root *Node) []int {
    // init result
    resultArr = make([]int, 0)
    nTraversal(root)
    return resultArr
}
func nTraversal(root *Node) {
    if root == nil {
        return
    }
    resultArr = append(resultArr, root.Val)

    for _, v := range root.Children {
        nTraversal(v)
    }
}
