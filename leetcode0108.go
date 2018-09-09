/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    return build(nums, 0, len(nums))
}

func build(nums []int, l, r int) *TreeNode {
    if l == r {
        return nil
    }
    mid := (l + r) / 2
    var node TreeNode
    node.Val = nums[mid]
    node.Left = sortedArrayToBST(nums[l:mid])
    node.Right = sortedArrayToBST(nums[mid + 1:r])
    return &node
}