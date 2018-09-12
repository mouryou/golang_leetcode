/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    l, r := minDepth(root.Left), minDepth(root.Right)
    if l < r {
        l, r = r, l
    }
    if r == 0 {
        return l + 1
    } else {
        return r + 1
    }
}