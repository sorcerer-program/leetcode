package linkedlist
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 // 相交链表
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
        return nil 
    }
    pA := headA
    pB := headB
    for ;; {
        if pA == pB {
            return pA
        } else if pA == nil {
            pA = headB
        } else if pB == nil {
            pB = headA
        } else {
            pA = pA.Next
            pB = pB.Next
        }
    }
}