package linkedlist

// leetcode707_DesignLinkedList


type MyLinkedList struct {
    dummyHead *ListNode
    size int
}

func Constructor() MyLinkedList {
    return MyLinkedList{&ListNode{}, 0}
}

func (this *MyLinkedList) Get(index int) int {
    if index < 0 || index >= this.size {
        return -1
    }
    p := this.dummyHead.Next
    for i := 0; i < index; i++ {
        p = p.Next
    }
    return p.Val
}

func (this *MyLinkedList) AddAtIndex(index, val int) {
    if index < 0 || index > this.size {
        return 
    }
    p := this.dummyHead
    for i := 0; i < index; i++ {
        p = p.Next
    }
    newNode := &ListNode{Val:val, Next:p.Next}
    p.Next = newNode
    this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
    if index < 0 || index >= this.size {
        return 
    }
    p := this.dummyHead
    for i := 0; i < index; i++ {
        p = p.Next
    }
    p.Next = p.Next.Next
    this.size--
}

func (this *MyLinkedList) AddAtHead(val int) {
    this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
    this.AddAtIndex(this.size, val)
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */