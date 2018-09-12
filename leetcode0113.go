/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
    results := make([][]int, 0)
    if root == nil {
        return results
    }
    search(root, []int {}, sum, &results)
    return results
}

func search(node *TreeNode, cur []int, sum int, results *[][]int) {
    cur = append(cur, node.Val)
    if node.Left == nil && node.Right == nil {
        if node.Val == sum {
            *results = append(*results, append([]int {}, cur...))
        }
    } else {
        if node.Left != nil {
            search(node.Left, cur, sum - node.Val, results)
        }
        if node.Right != nil {
            search(node.Right, cur, sum - node.Val, results)
        }
    }
}