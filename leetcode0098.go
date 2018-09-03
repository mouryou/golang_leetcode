/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    return validate(root, 0, 0, false, false)
}

func validate(node *TreeNode, l, r int, vl, vr bool) bool {
    if node == nil {
        return true
    }
    if vl {
        if node.Val <= l {
            return false
        }
    }
    if vr {
        if node.Val >= r {
            return false
        }
    }
    return validate(node.Left, l, node.Val, vl, true) && validate(node.Right, node.Val, r, true, vr)
}