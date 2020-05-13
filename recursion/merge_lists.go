/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    } else if l2 == nil {
        return l1
    }
    var head *ListNode
    var ptr *ListNode
    x := l1
    y := l2
    for x != nil || y != nil {
        if x == nil {
            ptr.Next = y
            ptr = ptr.Next
            y = y.Next
        } else if y == nil {
            ptr.Next = x
            ptr = ptr.Next
            x = x.Next
        } else if x.Val <= y.Val {
            if ptr != nil {
                ptr.Next = x
                ptr = ptr.Next
            } else {
                head = x
                ptr = head
            }
            x = x.Next
        } else {
            if ptr != nil {
                ptr.Next = y
                ptr = ptr.Next
            } else {
                head = y
                ptr = head
            }
            y = y.Next
        }
    }
    return head
}
