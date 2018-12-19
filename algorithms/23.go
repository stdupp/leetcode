/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    var min *ListNode
    var result *ListNode
    var exit bool
    var index int
    var last *ListNode
    
    for !exit {
        exit = true
        for i, list := range lists {
            if list != nil {
                exit = false
                if min != nil {
                    if list.Val < min.Val {
                        min = list
                        index = i
                    } 
                } else {
                    min = list
                    index = i
                }
            }
        }
        if min != nil {
            lists[index] = min.Next
        }
        
        if result == nil {
            last = min
            result = last
        } else {
            last.Next = min
            last = min
        }
        min = nil
    }
    return result
}
