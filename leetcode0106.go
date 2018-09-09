/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
    l := len(inorder)
    if l == 0 {
        return nil
    }
    val := postorder[l - 1]
    var node TreeNode
    node.Val = val
    i := 0
    for inorder[i] != val {
        i++
    }
    node.Left = buildTree(inorder[:i], postorder[:i])
    node.Right = buildTree(inorder[i + 1:], postorder[i:l - 1])
    return &node
}