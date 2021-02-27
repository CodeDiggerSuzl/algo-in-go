package tree

/*
给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明：叶子节点是指没有子节点的节点。
https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/
*/
func minDepthV1(root *TreeNode) int {
    if root == nil {
        return 0
    }
    if root.Left == nil {
        return minDepthV1(root.Right) + 1
    }
    if root.Right == nil {
        return minDepthV1(root.Left) + 1
    }
    ld := minDepthV1(root.Left)
    rd := minDepthV1(root.Right)
    if ld >= rd {
        return rd + 1
    } else {
        return ld + 1
    }
}

// better solution with only 2 recursion
func minDepthV2(root *TreeNode) int {
    if root == nil {
        return 0
    }

    ld := minDepthV2(root.Left)
    rd := minDepthV2(root.Right)
    if root.Left == nil || root.Right == nil {
        return ld + rd + 1
    }

    if ld >= rd {
        return rd + 1
    } else {
        return ld + 1
    }
}

/*
   public int minDepth_v1(TreeNode root) {
       if(root == null) return 0;
       int m1 = minDepth_v1(root.left);
       int m2 = minDepth_v1(root.right);
       //1.如果左孩子和右孩子有为空的情况，直接返回m1+m2+1
       //2.如果都不为空，返回较小深度+1
       return root.left == null || root.right == null ? m1 + m2 + 1 : Math.min(m1,m2) + 1;
   }
*/
