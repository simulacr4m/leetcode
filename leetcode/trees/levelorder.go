/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func traverse(root *TreeNode, queue []*TreeNode, list [][]int) [][]int {
    if root == nil {
        return list
    }
    queue = append(queue, root)
    list = append(list, []int{root.Val})
    for len(queue) > 0 {
        var new_queue []*TreeNode
        var sublist []int
        for _, pop := range queue {
            if pop.Left != nil {
                new_queue = append(new_queue, pop.Left)
                sublist = append(sublist, pop.Left.Val)
            }
            if pop.Right != nil {
                new_queue = append(new_queue, pop.Right)
                sublist = append(sublist, pop.Right.Val)
            }
        }
        queue = new_queue
        if len(sublist) > 0 {
            list = append(list, sublist)
        }
    }
    return list
}

func levelOrder(root *TreeNode) [][]int {
    var queue []*TreeNode
    var list [][]int
    list = traverse(root, queue, list)
    return list
}
