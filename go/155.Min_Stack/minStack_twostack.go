type MinStack struct {
    stack []int
    mins []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{}
}


func (this *MinStack) Push(x int)  {
    this.stack = append(this.stack, x)
    if len(this.mins) == 0 || x <= this.mins[len(this.mins)-1]{
        this.mins = append(this.mins, x)
    }
}


func (this *MinStack) Pop()  {
    top := this.stack[len(this.stack)-1]
    this.stack = this.stack[:len(this.stack)-1]
    if top <= this.mins[len(this.mins)-1] {
        this.mins = this.mins[:len(this.mins)-1]
    }
}


func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
    return this.mins[len(this.mins)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
