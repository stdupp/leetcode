/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var result [][]int

func pathSum(root *TreeNode, sum int) [][]int {
    if root == nil {
        return nil
    }
    result = make([][]int, 0, 1)
    find(root, []int{}, sum)
    return result
}

func find(root *TreeNode, nums []int, sum int ) {
    nums2 := make([]int, len(nums) + 1)
    copy(nums2, nums)
    nums2[len(nums)] = root.Val
    if root.Left == nil && root.Right == nil {
        if root.Val == sum {
            result = append(result, nums2)
        }
        return
    }
    if root.Left != nil {
        find(root.Left, nums2, sum - root.Val)
    }
    if root.Right != nil {
        find(root.Right, nums2, sum - root.Val)
    }
} 
