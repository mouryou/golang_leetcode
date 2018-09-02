/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
    if m == n {
        return head
    }
    var dummy ListNode
    dummy.Next = head
    n1 := &dummy
    for i := 0; i < m - 1; i++ {
        n1 = n1.Next
    }
    n2, n3 := n1.Next, n1.Next.Next
    for i := 0; i < n - m; i++ {
        n4 := n3.Next
        n3.Next = n2
        n2, n3 = n3, n4
    }
    n1.Next.Next = n3
    n1.Next = n2
    return dummy.Next
}