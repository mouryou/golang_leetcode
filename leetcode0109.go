/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
    nums := make([]int, 0)
    for n := head; n != nil; n = n.Next {
        nums = append(nums, n.Val)
    }
    return build(nums, 0, len(nums))
}

func build(nums []int, l, r int) *TreeNode {
    if l == r {
        return nil
    }
    mid := (l + r) / 2
    return &TreeNode {
        Val: nums[mid],
        Left: build(nums, l, mid),
        Right: build(nums, mid + 1, r),
    }
}