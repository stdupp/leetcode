/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
    gap := n - m
    if gap == 0 {
        return head
    }
    var pre, next *ListNode
    node1, node2 := head, head
    for gap > 0 {
        node1 = node1.Next
        next = node1.Next
        gap--
    }
    
    for i := m -1; i > 0; i-- {
        pre = node2
        node1 = node1.Next
        node2 = node2.Next
        next = node1.Next
    }
    tail := next
    for node2 != node1 {
        tmp := tail
        tail = node2
        node2 = node2.Next
        tail.Next = tmp
    }
    node1.Next = tail
    if pre != nil {
        pre.Next = node1
        return head
    }
    
    return node1
}
