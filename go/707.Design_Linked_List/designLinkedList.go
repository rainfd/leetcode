type MyLinkedList struct {
    Val int
    Next *MyLinkedList
}


/** Initialize your data structure here. */
func Constructor() MyLinkedList {
    // head node
    return MyLinkedList{}
}


/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
    node := this
    for i := 0; i <= index; i++ {
        if node.Next == nil {
            break
        }
        if index == i {
            return node.Next.Val
        }
        node = node.Next
    }
    return -1
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
    newNode := &MyLinkedList{
        Val: val,
        Next: this.Next,
    }
    this.Next = newNode
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
    node := this
    for node.Next != nil {
        node = node.Next
    }
    node.Next = &MyLinkedList{
        Val: val,
        Next: nil,
    }
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    node := this
    for i := 0; i <= index; i++ {
        if node.Next == nil && i != index {
            return
        }
        if i == index {
            node.Next = &MyLinkedList{
                Val: val,
                Next: node.Next,
            }
            return
        }
        node = node.Next
    }
    
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
    node := this
    for i := 0; i <= index; i++ {
        if node.Next == nil {
            return
        }
        if i == index {
            node.Next = node.Next.Next
        }
        node = node.Next
    }
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
