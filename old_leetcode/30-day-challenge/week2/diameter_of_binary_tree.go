/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

func height(node *TreeNode) int {
    if node == nil {
        return 0
    }
    lheight := height(node.Left)
    rheight := height(node.Right)
    return 1 + max(lheight, rheight)
}

func diameter(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    ldiameter := diameter(root.Left)
    rdiameter := diameter(root.Right)
    
    return max(1 + height(root.Left) + height(root.Right), max(ldiameter, rdiameter))
}

func diameterOfBinaryTree(root *TreeNode) int {
    d := diameter(root)
    if d == 0 {
        return 0
    }
    return d-1
}
