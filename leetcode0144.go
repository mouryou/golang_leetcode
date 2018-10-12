/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    result := []int {}
    if root == nil {
        return result
    }
    stack := []*TreeNode {root}
    for len(stack) > 0 {
        n := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        result = append(result, n.Val)
        if n.Right != nil {
            stack = append(stack, n.Right)
        }
        if n.Left != nil {
            stack = append(stack, n.Left)
        }
    }
    return result
}