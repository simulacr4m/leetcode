func removeElements(head *ListNode, val int) *ListNode {
    var prev *ListNode
    x := head
    for x != nil {
        if x.Val == val {
            before := prev
            after := x.Next
            if before != nil {
                before.Next = after
            } else {
                head = after
            }
        } else {
            prev = x
        }
        x = x.Next
    }
    return head
}
