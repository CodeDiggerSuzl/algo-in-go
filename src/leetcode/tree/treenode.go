package tree

// Node def for bin tree
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// Node of n-ary tree
type NNode struct {
    Val      int
    Children []*NNode
}
