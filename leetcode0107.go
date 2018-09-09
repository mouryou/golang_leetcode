/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
    result := make([][]int, 0)
    search(root, 0, &result)
    return result
}

func search(node *TreeNode, level int, result *[][]int) {
    if node == nil {
        return
    }
    l := len(*result)
    if level == l {
        *result = append([][]int {[]int {node.Val}}, (*result)...)
    } else {
        (*result)[l - 1 - level] = append((*result)[l - 1 - level], node.Val)
    }
    search(node.Left, level + 1, result)
    search(node.Right, level + 1, result)
}