/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func swapHelper(first *ListNode) *ListNode {
    if first == nil {
        return nil
    }
    
    second := first.Next
    if second == nil {
        return first
    }
    after := second.Next
    
    second.Next = first
    first.Next = after
    first.Next = swapHelper(after)
    return second
}

func swapPairs(head *ListNode) *ListNode {
    return swapHelper(head)
}
