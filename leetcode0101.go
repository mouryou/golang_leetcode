/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return compare(root.Left, root.Right)
}

func compare(n1, n2 *TreeNode) bool {
    if n1 == nil && n2 == nil {
        return true
    } else if n1 == nil || n2 == nil {
        return false
    } else {
        return n1.Val == n2.Val && compare(n1.Left, n2.Right) && compare(n1.Right, n2.Left)
    }
}