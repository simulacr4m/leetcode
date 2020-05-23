/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    var list, last *ListNode
    x, y := l1, l2
    for x != nil && y != nil {
        var node, after *ListNode
        if x.Val <= y.Val {
            node = x
        } else {
            node = y
        }
        if last != nil {
            after = node.Next
            node.Next = nil
            last.Next = node
            last = last.Next
        } else {
            after = node.Next
            node.Next = nil
            list = node
            last = list
        }
        if node == x {
            x = after
        } else {
            y = after
        }
    }
    if x != nil {
        if last != nil {
            last.Next = x
        } else {
            list = x
        }
    }
    if y != nil {
        if last != nil {
            last.Next = y
        } else {
            list = y
        }
    }
    return list
}
