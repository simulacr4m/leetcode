type Node struct {
    val int
    next *Node
}

type MinStack struct {
    first *Node
    mins *Node
}

func insertMins(this *MinStack, x int) {
    if this.mins == nil {
        this.mins = &Node{val: x, next: nil}
        return
    }
    
    var prev *Node
    var n *Node = this.mins
    for ; n != nil; n = n.next {
        if x <= n.val {
            xN := &Node{val: x, next: n}
            if prev != nil {
                prev.next = xN
            } else {
                this.mins = xN
            }
            return
        }
        prev = n
    }
    if n == nil {
        prev.next = &Node{val: x, next: nil}
    }
}

func removeMins(this *MinStack, x int) {
    var prev *Node
    for n := this.mins; n != nil; n = n.next {
        if n.val == x {
            if prev != nil {
                prev.next = n.next
            } else {
                this.mins = this.mins.next
            }
            fmt.Printf("\n");
            return
        }
        prev = n
    }
}

func returnTopMin(this *MinStack) int {
    if this.mins == nil {
        return 0
    }
    return this.mins.val
}

/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{}
}


func (this *MinStack) Push(x int)  {
    n := &Node{val: x, next: nil}
    n.next = this.first
    this.first = n
    insertMins(this, x)
}


func (this *MinStack) Pop()  {
    if this.first == nil {
        return
    }
    removeMins(this, this.first.val)
    this.first = this.first.next
}


func (this *MinStack) Top() int {
    if this.first == nil {
        return 0
    }
    return this.first.val
}


func (this *MinStack) GetMin() int {
    return returnTopMin(this)
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
