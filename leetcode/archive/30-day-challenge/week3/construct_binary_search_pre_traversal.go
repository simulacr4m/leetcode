/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func bstBuilder(preorder []int, lo, hi int) *TreeNode {
    if lo < 0 || lo > hi {
        return nil
    }
    root, i := &TreeNode{Val: preorder[lo]}, lo+1
    for ; i <= hi; i++ {
        if preorder[i] > root.Val {
            break
        }
    }
    root.Left = bstBuilder(preorder, lo+1, i-1)
    root.Right = bstBuilder(preorder, i, hi)
    return root
}

func bstFromPreorder(preorder []int) *TreeNode {
    return bstBuilder(preorder, 0, len(preorder)-1)
}
