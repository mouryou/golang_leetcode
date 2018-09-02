/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    res := []int {}
    stack := []*TreeNode {root}
    for len(stack) > 0 {
        n := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        if n == nil {
            continue
        } 
        if len(stack) > 0 && stack[len(stack) - 1] == n.Right {
            res = append(res, n.Val)
        } else {
            stack = append(stack, n.Right)
            stack = append(stack, n)
            stack = append(stack, n.Left)
        }
    }
    return res
}