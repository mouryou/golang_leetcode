/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
    if root == nil {
        return 0
    }
    res := 0
    sum(root, 0, &res)
    return res
}

func sum(node *TreeNode, curSum int, res *int) {
    curSum = curSum * 10 + node.Val
    if node.Left == nil && node.Right == nil {
        *res += curSum
        return
    }
    if node.Left != nil {
        sum(node.Left, curSum, res)
    }
    if node.Right != nil {
        sum(node.Right, curSum, res)
    }
}