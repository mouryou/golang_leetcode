/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    a, b, c := head, head.Next, head.Next.Next
    a.Next = swapPairs(c)
    b.Next = a
    return b
}