/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    l := 1
    n := head
    for n.Next != nil {
        n = n.Next
        l++
    }
    k = k % l
    if k == 0 {
        return head
    }
    n.Next = head
    for i := 0; i < l - k - 1; i++ {
        head = head.Next
    }
    res := head.Next
    head.Next = nil
    return res
}