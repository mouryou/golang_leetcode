/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    n := head
    for n.Next != nil {
        if n.Next.Val == n.Val {
            n.Next = n.Next.Next
        } else {
            n = n.Next
        }
    }
    return head
}