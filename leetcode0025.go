/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    if k == 1 || head == nil {
        return head
    }
    p := head
    for i := 0; i < k - 1; i++ {
        p = p.Next
        if p == nil {
            return head
        }
    }
    tail := reverseKGroup(p.Next, k)
    p1, p2 := head, head.Next
    for p1 != p {
        t := p2.Next
        p2.Next = p1
        p1 = p2
        p2 = t
    }
    head.Next = tail
    return p
}