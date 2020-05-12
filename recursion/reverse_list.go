/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


func reverseListHelper(head *ListNode, first **ListNode) {
    if head == nil || head.Next == nil {
        *first = head
        return
    }
    reverseListHelper(head.Next, first)
    head.Next.Next = head
    head.Next = nil
}

func reverseList(head *ListNode) *ListNode {
    var first *ListNode
    reverseListHelper(head, &first)
    return first
}
