/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    l := remove(head, n)
    if l + 1 == n {
        return head.Next
    } else {
        return head
    }
}

func remove(node *ListNode, n int) int {
    var i int
    if node.Next == nil {
        i = 0
    } else {
        i = remove(node.Next, n) + 1
    }
    if i == n {
        node.Next = node.Next.Next
    }
    return i
}
