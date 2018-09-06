/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode)  {
    var last, prev, cur, n1, n2 *TreeNode
    cur = root
    find := false
    for cur != nil {
        if cur.Left == nil {
            if last != nil && cur != nil && cur.Val < last.Val {
                n2 = cur
                if !find {
                    n1 = last
                    find = true
                }
            }
            last = cur
            cur = cur.Right
        } else {
            prev = cur.Left
            for prev.Right != nil && prev.Right != cur {
                prev = prev.Right
            }
            if prev.Right == nil {
                prev.Right = cur
                cur = cur.Left
            } else {
                prev.Right = nil
                if last != nil && cur != nil && cur.Val < last.Val {
                    n2 = cur
                    if !find {
                        n1 = last
                        find = true
                    }
                }
                last = cur
                cur = cur.Right
            }
        }
    }
    n1.Val, n2.Val = n2.Val, n1.Val
}