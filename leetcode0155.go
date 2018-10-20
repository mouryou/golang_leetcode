type MinStack struct {
    stack []int
    mins []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{
        stack: []int {},
        mins: []int {},
    }
}


func (this *MinStack) Push(x int)  {
    l := len(this.stack)
    this.stack = append(this.stack, x)
    if l == 0 || x < this.mins[l - 1] {
        this.mins = append(this.mins, x)
    } else {
        this.mins = append(this.mins, this.mins[l - 1])
    }
}


func (this *MinStack) Pop()  {
    l := len(this.stack)
    if l > 0 {
        this.stack = this.stack[:l - 1]
        this.mins = this.mins[:l - 1]
    }
}


func (this *MinStack) Top() int {
    l := len(this.stack)
    if l == 0 {
        return 0
    }
    return this.stack[l - 1]
}


func (this *MinStack) GetMin() int {
    l := len(this.mins)
    if l == 0 {
        return 0
    }
    return this.mins[l - 1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */