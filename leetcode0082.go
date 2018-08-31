/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    var dumb ListNode
    dumb.Next = head
    n := &dumb
    for n.Next != nil && n.Next.Next != nil {
        if n.Next.Next.Val == n.Next.Val {
            t := n.Next.Next.Next
            for t != nil && t.Val == n.Next.Val {
                t = t.Next
            }
            n.Next = t
        } else {
            n = n.Next
        }
    }
    return dumb.Next
}