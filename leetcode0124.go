/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    result := root.Val
    search(root, &result)
    return result
}

func search(node *TreeNode, result *int) int {
    if node == nil {
        return 0
    }
    l, r := search(node.Left, result), search(node.Right, result)
    t := node.Val
    if l > 0 {
        t += l
    }
    if r > 0 {
        t += r
    }
    if t > *result {
        *result = t
    }
    t = node.Val
    if l < r {
        l, r = r, l
    }
    if l > 0 {
        t += l
    }
    if t > 0 {
        return t
    } else {
        return 0
    }
}