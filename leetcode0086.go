/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    var less, more ListNode
    nl, nm := &less, &more
    for head != nil {
        if head.Val < x {
            nl.Next = head
            nl = nl.Next
        } else {
            nm.Next = head
            nm = nm.Next
        }
        head = head.Next
    }
    nl.Next = more.Next
    nm.Next = nil
    return less.Next
}