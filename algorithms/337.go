/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
    r := dfs(root)
    if r[0] > r[1] {
        return r[0]
    }  
    return r[1]
}

func dfs(root *TreeNode) [2]int {
    if root == nil {
        return [2]int{0,0}
    }

    left := dfs(root.Left)
    right := dfs(root.Right)
    
    r1 := root.Val + left[1] + right[1]
    r2 := 0
    if left[0] > left[1] {
        r2 += left[0]
    } else {
        r2 += left[1]
    }
    if right[0] > right[1] {
        r2 += right[0]
    } else {
        r2 += right[1]
    }
    return [2]int{r1, r2}
}
