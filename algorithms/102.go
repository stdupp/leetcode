/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    var result [][]int
    stack := []*TreeNode{root}
    
    for len(stack) > 0 {
        var tmp []int
        var levelNode []*TreeNode
        for _, node := range stack {
            if node == nil {
                continue
            }
            tmp = append(tmp, node.Val)
            levelNode = append(levelNode, node.Left, node.Right)
        }
        if len(tmp) > 0 {
            result = append(result, tmp)   
        }
        stack = levelNode
    }
    return result
}
