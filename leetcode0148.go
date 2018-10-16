/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    slow, fast := head, head.Next
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    head2 := slow.Next
    slow.Next = nil
    head1, head2 := sortList(head), sortList(head2)
    var n ListNode
    p := &n
    for head1 != nil && head2 != nil {
        if head1.Val > head2.Val {
            head1, head2 = head2, head1
        }
        p.Next = head1
        p, head1 = p.Next, head1.Next
    }
    if head1 != nil {
        p.Next = head1
    }
    if head2 != nil {
        p.Next = head2
    }
    return n.Next
}