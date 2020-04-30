/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func max(a int, b int) int {
    if a >= b {
        return a
    }
    return b
}

func maxPathHelper(root *TreeNode, res *int) int {
    if root == nil {
        return 0
    }
    
    leftSum := maxPathHelper(root.Left, res)
    rightSum := maxPathHelper(root.Right, res)
    
    maxSingle := max(max(leftSum, rightSum) + root.Val, root.Val)
    maxTop := max(maxSingle, leftSum + rightSum + root.Val)
    *res = max(*res, maxTop)
    return maxSingle
}

func maxPathSum(root *TreeNode) int {
    res := math.MinInt32
    maxPathHelper(root, &res)
    return res
}
