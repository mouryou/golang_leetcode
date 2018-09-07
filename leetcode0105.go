/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    l := len(preorder)
    if l == 0 {
        return nil
    }
    i := 0
    for inorder[i] != preorder[0] {
        i++
    }
    var root TreeNode
    root.Val = preorder[0]
    root.Left = buildTree(preorder[1:i + 1], inorder[:i])
    root.Right = buildTree(preorder[i + 1:], inorder[i + 1:])
    return &root
}