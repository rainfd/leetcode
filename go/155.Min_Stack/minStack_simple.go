type MinStack struct {
    stack []int
    min int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{}
}


func (this *MinStack) Push(x int) {
    if len(this.stack) == 0 {
        this.min = x
    } else if x < this.min {
        this.min = x
    }
    this.stack = append(this.stack, x)
}


func (this *MinStack) Pop()  {
    length := len(this.stack)
    this.stack = this.stack[:length-1]
    if len(this.stack) == 0 {
        this.min = 0
        return
    }
    min := this.stack[0]
    for _, n := range this.stack {
        if n < min {
            min = n
        }
    }
    this.min = min
}


func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
    return this.min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
