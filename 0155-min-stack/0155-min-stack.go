type MinStack struct {
    stack []Node
}

type Node struct {
    val int
    min int
}

func Constructor() MinStack {
    return MinStack{stack: []Node{}}
}


func (this *MinStack) Push(val int)  {
    newMin := val
    // If stack isn't empty, compare new val with current min
    if len(this.stack) > 0 {
        currentMin := this.stack[len(this.stack)-1].min
        if currentMin < newMin {
            newMin = currentMin
        }
    }
    // Push the pair onto the stack
    this.stack = append(this.stack, Node{val: val, min: newMin})
}


func (this *MinStack) Pop()  {
    if len(this.stack) > 0 {
        this.stack = this.stack[:len(this.stack)-1]
    }
}


func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1].val
}


func (this *MinStack) GetMin() int {
    return this.stack[len(this.stack)-1].min
}
