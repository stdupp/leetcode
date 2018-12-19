/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    tmp := 0
    var result *ListNode
    var tail *ListNode

    if l1 == nil {
        return l2
    }

    for l1 != nil {
        v1 := l1.Val
        v2 := 0
        if l2 != nil {
            v2 = l2.Val
            l2 = l2.Next
        }
        v := v1+v2+tmp
        if v >= 10 {
            v = v- 10
            tmp = 1
        } else {
            tmp = 0
        }
        if result == nil {
            result = &ListNode{v, nil}
            tail = result
        } else {
            tail.Next = &ListNode{v, nil}
            tail = tail.Next
        }
        l1 = l1.Next
    }
    
    for l2 != nil {
        v := l2.Val + tmp
        if v >= 10 {
            v = v - 10
            tmp = 1
        } else {
            tmp = 0
        }
        if result == nil {
            result = &ListNode{v, nil}
            tail = result
        } else {
            tail.Next = &ListNode{v, nil}
            tail = tail.Next
        }
        l2 = l2.Next
    }
    
    if tmp != 0 {
        tail.Next = &ListNode{1, nil}
    }
    
    return result
    
}
