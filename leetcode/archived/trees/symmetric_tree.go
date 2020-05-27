/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func traverseHelper(root *TreeNode, list []int, left bool) []int {
    if root == nil {
        list = append(list, -1)
        return list
    }
    list = append(list, root.Val)
    if left {
        list = traverseHelper(root.Left, list, left)
        list = traverseHelper(root.Right, list, left)
    } else {
        list = traverseHelper(root.Right, list, left)
        list = traverseHelper(root.Left, list, left)
    }
    return list
}

func traverse(root *TreeNode, left bool) []int {
    var list []int
    list = traverseHelper(root, list, left)
    return list
}

func isSymmetric(root *TreeNode) bool {
    left, right := traverse(root, true), traverse(root, false)
    if len(left) != len(right) {
        return false
    }
    for i := 0; i < len(left); i++ {
        if left[i] != right[i] {
            return false
        }
    }
    return true
}
