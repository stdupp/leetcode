/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
    if n < 1 {
        return nil
    }
    
    return geneTree(1, n)
}

func geneTree(start, end int)[]*TreeNode {
    if start > end {
        return []*TreeNode{nil}
    }
    
    var r []*TreeNode
    for i := start; i <= end; i++ {
        left := geneTree(start, i - 1)
        right := geneTree(i+1, end)
        for _, vl := range left {
            for _, vr := range right {
                t := &TreeNode{i, vl, vr}
                r = append(r, t)
            }
        }
    }
    return r
}
