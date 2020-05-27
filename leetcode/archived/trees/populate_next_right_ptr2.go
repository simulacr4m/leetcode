/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
        return nil
    }
    var queue []*Node = []*Node{root}
    for len(queue) > 0 {
        var level []*Node
        for _, pop := range queue {
            if pop == nil {
                continue
            }
            if pop.Left != nil {
                level = append(level, pop.Left)
            }
            if pop.Right != nil {
                level = append(level, pop.Right)
            }
        }
        for i := 0; i < len(level)-1; i++ {
            level[i].Next = level[i+1]
        }
        queue = level
    }
    return root
}
