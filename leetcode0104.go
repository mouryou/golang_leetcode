/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    result := maxDepth(root.Left)
    t := maxDepth(root.Right)
    if t > result {
        result = t
    }
    return result + 1
}