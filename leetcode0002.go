/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    res := &ListNode{}
    t := res
    c := 0
    for l1 != nil || l2 != nil || c != 0 {
        t.Next = &ListNode{}
        t = t.Next
        t.Val += c
        if l1 != nil {
            t.Val += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            t.Val += l2.Val
            l2 = l2.Next
        }
        c = t.Val / 10
        t.Val %= 10
    }
    return res.Next
}