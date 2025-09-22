package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummyHead := &ListNode{0, head}
    left := dummyHead
    right := dummyHead
    for i := 0; i < n; i++{
        right = right.Next
    }
    for ;right.Next != nil; {
        left = left.Next
        right = right.Next
    }
    left.Next = left.Next.Next
    return dummyHead.Next
}