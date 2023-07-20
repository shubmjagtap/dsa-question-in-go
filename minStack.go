type MinStack struct {
    stack [][]int
}

func Constructor() MinStack {
    return MinStack{}
}

func (this *MinStack) Push(val int) {
    if len(this.stack) == 0 {
        // If the stack is empty, store the current value as the minimum value.
        this.stack = append(this.stack, []int{val, val})
    } else {
        // Calculate the new minimum value based on the current value and the previous minimum.
        prevMin := this.stack[len(this.stack)-1][1]
        newMin := val
        if val > prevMin {
            newMin = prevMin
        }
        this.stack = append(this.stack, []int{val, newMin})
    }
}

func (this *MinStack) Pop() {
    if len(this.stack) > 0 {
        this.stack = this.stack[:len(this.stack)-1]
    }
}

func (this *MinStack) Top() int {
    if len(this.stack) > 0 {
        return this.stack[len(this.stack)-1][0]
    }
    return -1
}

func (this *MinStack) GetMin() int {
    if len(this.stack) > 0 {
        return this.stack[len(this.stack)-1][1]
    }
    return -1
}



/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
