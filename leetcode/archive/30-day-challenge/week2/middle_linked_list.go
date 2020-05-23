func middleNode(head *ListNode) *ListNode {
    var list []*ListNode
    ptr := head
    for ptr != nil {
        list = append(list, ptr)
        ptr = ptr.Next
    }
    return list[len(list)/2]
}
