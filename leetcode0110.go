/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    return check(root) != -1
}

func check(node *TreeNode) int {
    if node == nil {
        return 0
    }
    l, r := check(node.Left), check(node.Right)
    if l < r {
        l, r = r, l
    }
    if r == -1 {
        return -1
    }
    if l - r > 1 {
        return -1
    } else {
        return l + 1
    }
}