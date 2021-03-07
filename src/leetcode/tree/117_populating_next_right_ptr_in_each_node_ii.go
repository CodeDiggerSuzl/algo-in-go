package tree

/*
给定一个 二叉树(不是完美二叉树)!!! ，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：

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

// var ru_ []int
// 想用下面的方法,但是一直想用[]*Node 来寸,但是没法判断  if len(ru_)== depth && ru_[depth] !=nil
// 总会出现空指针,应该用 map 的 !
func connectNil(root *NodeWithNext) *NodeWithNext {
	maps := make(map[int]*NodeWithNext)
	maps = dfsNil(root, 0, maps)
	return root
}
func dfsNil(root *NodeWithNext, depth int, maps map[int]*NodeWithNext) map[int]*NodeWithNext {
	if root == nil {
		return maps
	}
	dfsNil(root.Right, depth+1, maps)
	dfsNil(root.Left, depth+1, maps)
	if _, ok := maps[depth]; ok {
		root.Next = maps[depth]
	} else {
		root.Next = nil
	}
	maps[depth] = root
	return maps
}

//链接：https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/solution/17-tian-chong-mei-ge-jie-dian-de-xia-yi-ge-you-ce-/
