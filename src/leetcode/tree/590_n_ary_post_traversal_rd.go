package tree

/*
给定一个 N 叉树，返回其节点值的后序遍历。

例如，给定一个 3叉树 :

        1
      / | \
    3  2   4
   / \
  5   6

返回其后序遍历: [5,6,3,2,4,1].

说明: 递归法很简单，你可以使用迭代法完成此题吗?
*/

// 递归方法
var nRes []int

func nPostorder(root *Node) []int {
    nRes = make([]int, 0)
    nPostTraversal(root)
    return nRes
}

func nPostTraversal(root *Node) {
    if root == nil {
        return
    }
    for _, node := range root.Children {
        nPostTraversal(node)
    }
    nRes = append(nRes, root.Val)
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
// 使用迭代方法进行解决
func nPostorderIter(root *Node) []int {
    res := make([]int, 0)
    if root == nil {
        return []int{}
    }

    return res
}
