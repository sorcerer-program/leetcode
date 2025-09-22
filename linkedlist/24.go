package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 24. Swap Nodes in Pairs
func swapPairs(head *ListNode) *ListNode {
    dummy_head := &ListNode{0, head}
    pre := dummy_head
    for ; pre.Next !=nil && pre.Next.Next != nil; {
        p := pre.Next
        next := p.Next
        temp := next.Next
        pre.Next = next
        next.Next = p
        p.Next = temp
        pre = p
    }
    return dummy_head.Next
}