package tree

/*
给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

初始状态下，所有next 指针都被设置为 NULL。
链接：https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node
*/

type NodeWithNext struct {
	Val   int
	Left  *NodeWithNext
	Right *NodeWithNext
	Next  *NodeWithNext
}

// 层序遍历 方法以

func connect(root *NodeWithNext) *NodeWithNext {
	if root == nil {
		return nil
	}
	connectEach(root.Left, root.Right)
	return root
}

func connectEach(n1 *NodeWithNext, n2 *NodeWithNext) {
	if n1 == nil || n2 == nil {
		return
	}
	n1.Next = n2
	// 相同父节点的 node
	connectEach(n1.Left, n1.Right)
	connectEach(n2.Left, n2.Right)
	// 不同父节点的 node
	connectEach(n1.Right, n2.Left)
}