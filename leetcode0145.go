/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    res := []int {}
    if root == nil {
        return res
    }
    stack := []*TreeNode {root}
    for len(stack) > 0 {
        n := stack[len(stack) - 1]
        if n.Left == nil && n.Right == nil {
            res = append(res, n.Val)
            stack = stack[:len(stack) - 1]
        } else {
            if n.Right != nil {
                stack = append(stack, n.Right)
                n.Right = nil
            }
            if n.Left != nil {
                stack = append(stack, n.Left)
                n.Left = nil
            }
        }
    }
    return res
}