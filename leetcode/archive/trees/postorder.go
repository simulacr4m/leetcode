/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func traverse(root *TreeNode, list []int) []int {
    if root == nil {
        return list
    }
    list = traverse(root.Left, list)
    list = traverse(root.Right, list)
    list = append(list, root.Val)
    return list
}

func postorderTraversal(root *TreeNode) []int {
    var list []int
    return traverse(root, list)
}
