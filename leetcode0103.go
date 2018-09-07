/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    result := make([][]int, 0)
    search(root, 0, &result)
    return result
}

func search(node *TreeNode, level int, result *[][]int) {
    if node == nil {
        return
    }
    if level == len(*result) {
        *result = append(*result, []int {node.Val})
    } else if level % 2 == 0 {
        (*result)[level] = append((*result)[level], node.Val)
    } else {
        (*result)[level] = append([]int {node.Val}, (*result)[level]...)
    }
    search(node.Left, level + 1, result)
    search(node.Right, level + 1, result)
}