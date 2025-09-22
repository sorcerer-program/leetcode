package linkedlist

// 移出链表元素
type ListNode struct {
	Val int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
    dummyHead := &ListNode{Val:0, Next:head}
    for p := dummyHead; p.Next != nil;  {
        if (p.Next.Val == val) {
            p.Next = p.Next.Next
        }else {
            p = p.Next
        }
    }
    return dummyHead.Next
}
