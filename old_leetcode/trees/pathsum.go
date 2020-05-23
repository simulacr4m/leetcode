func PathSum(root *TreeNode, target int) int {
    if root == nil {
        return 0
    }
    
    left, right := target-1, target-1
    
    if root.Left != nil {
        left = root.Val + PathSum(root.Left, target - root.Val)
    }
    if root.Right != nil {
        right = root.Val + PathSum(root.Right, target - root.Val)
    }
    if root.Left == nil && root.Right == nil {
        return root.Val
    }
    if left == target || right == target {
        return target
    }
    return left
}

func hasPathSum(root *TreeNode, sum int) bool {
    if root == nil {
        return false
    }
    res := PathSum(root, sum)
    return res == sum
}
