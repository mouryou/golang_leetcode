/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func upsideDownBinaryTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    a, b, c := root, root.Left, root.Right
    root.Left, root.Right = nil, nil
    for b != nil {
        b.Left, b.Right, a, b, c = c, a, b, b.Left, b.Right
    }
    return a
}