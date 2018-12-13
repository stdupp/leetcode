/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    var pre *ListNode
    var next *ListNode
    var nextCur *ListNode
    
    r := head.Next
    current := head
    
    for current != nil && current.Next != nil {
        next = current.Next
        nextCur = next.Next
        next.Next = current
        current.Next = nextCur
        if pre != nil {
            pre.Next = next
        }
        pre = current
        current = nextCur
    }
    
    return r
}
