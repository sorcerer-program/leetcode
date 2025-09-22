package linkedlist

func reverseList(head *ListNode) *ListNode {
    var pre *ListNode = nil
    for cur := head; cur != nil; {
        behind := cur.Next
        cur.Next = pre
        pre = cur
        cur = behind
    }
    return pre
}