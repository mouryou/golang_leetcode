/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
    if head == nil || head.Next == nil {
        return
    }
    mid, end := head, head.Next
    for end != nil && end.Next != nil {
        mid = mid.Next
        end = end.Next.Next
    }
    head2 := mid.Next
    mid.Next = nil
    head2 = reverse(head2)
    p1, p2 := head, head2
    for p2 != nil {
        t1, t2 := p1.Next, p2.Next
        p1.Next = p2
        p2.Next = t1
        p1, p2 = t1, t2
    }
}

func reverse(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    a, b := head, head.Next
    a.Next = nil
    for b != nil {
        t := b.Next
        b.Next = a
        a, b = b, t
    }
    return a
}