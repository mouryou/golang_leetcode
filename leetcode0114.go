/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
    search(root)
}

func search(node *TreeNode) (head, tail *TreeNode) {
    if node == nil {
        return nil, nil
    }
    head, tail = node, node
    lh, lt := search(node.Left)
    rh, rt := search(node.Right)
    if lh != nil {
        tail.Right = lh
        tail = lt
    }
    if rh != nil {
        tail.Right = rh
        tail = rt
    }
    head.Left = nil
    return head, tail
}