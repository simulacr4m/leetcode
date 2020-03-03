/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

import "fmt"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var res, ptr *ListNode
    x, y := l1, l2
    carry := false
    for x != nil && y != nil {
        a, b := x.Val, y.Val
        sum := a + b
        if carry {
            sum += 1
            carry = false
        }
        if sum > 9 {
            sum -= 10
            carry = true
        }
        if res == nil {
            res = &ListNode{sum, nil}
            ptr = res
        } else {
            ptr.Next = &ListNode{sum, nil}
            ptr = ptr.Next
        }
        x = x.Next
        y = y.Next
    }
    for x != nil {
        if carry {
            x.Val += 1
            carry = false
        }
        if x.Val > 9 {
            x.Val -= 10
            carry = true
        }
        ptr.Next = &ListNode{x.Val, nil}
        ptr = ptr.Next
        x = x.Next
    }
    for y != nil {
        if carry {
            y.Val += 1
            carry = false
        }
        if y.Val > 9 {
            y.Val -= 10
            carry = true
        }
        ptr.Next = &ListNode{y.Val, nil}
        ptr = ptr.Next
        y = y.Next
    }
    if carry {
        ptr.Next = &ListNode{1, nil}
    }
    return res
}
