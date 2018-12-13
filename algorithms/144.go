/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }

    var stack []*TreeNode
    var result []int
    node := root
    for {
        result = append(result, node.Val)
        stack = append(stack, node)
        if node.Left != nil {
            node = node.Left
        } else {
            node = stack[len(stack) -1]
            stack = stack[:len(stack)-1]
            for node.Right == nil && len(stack) > 0 {
                node = stack[len(stack) -1]
                stack = stack[:len(stack)-1]
            }
            if node.Right != nil {
                node = node.Right   
            } else {
                break
            }
        }
    }
    return result
}
