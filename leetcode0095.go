/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
    if n == 0 {
        return []*TreeNode {}
    }
    return generate(1, n + 1, make(map[[2]int][]*TreeNode))
}

func generate(l, r int, dp map[[2]int][]*TreeNode) []*TreeNode {
    if l == r {
        return []*TreeNode {nil}
    }
    if r, ok := dp[[2]int {l, r}]; ok {
        return r
    }
    res := []*TreeNode {}
    for i := l; i < r; i++ {
        for _, left := range generate(l, i, dp) {
            for _, right := range generate(i + 1, r, dp) {
                var node TreeNode
                node.Val = i
                node.Left = left
                node.Right = right
                res = append(res, &node)
            }
        }
    }
    dp[[2]int {l, r}] = res
    return res
}