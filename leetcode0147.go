/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
    var res ListNode
    p := &res
    for head != nil {
        if !(p.Next != nil && p.Next.Val < head.Val) {
            p = &res
        }
        for p.Next != nil && p.Next.Val < head.Val {
            p = p.Next
        }
        p.Next, head.Next, head = head, p.Next, head.Next
    }
    return res.Next
}