/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }
    l := len(lists) // length of current lists
    for l > 1 {
        for i := 0; i < l / 2; i++ {
            lists[i] = merge2Lists(lists[i], lists[l - i - 1])
        } 
        l = (l + 1) / 2
    }
    return lists[0]
}

func merge2Lists(l1, l2 *ListNode) *ListNode {
    var head ListNode
    p := &head
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            p.Next = l1
            l1 = l1.Next
        } else {
            p.Next = l2
            l2 = l2.Next
        }
        p = p.Next
    }
    if l1 != nil {
        p.Next = l1
    }
    if l2 != nil {
        p.Next = l2
    }
    return head.Next
}