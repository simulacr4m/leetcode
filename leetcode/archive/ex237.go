/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
    var prev *ListNode
    for node.Next != nil {
        node.Val = node.Next.Val
        prev = node
        node = node.Next
    }
    if prev != nil{
        prev.Next = nil
    }
}
