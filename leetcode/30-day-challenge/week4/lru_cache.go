type Node struct {
    Key int
    Value int
    Prev *Node
    Next *Node
}

type LRUCache struct {
    n int
    capacity int
    dict map[int]*Node
    head *Node
    tail *Node
}


func Constructor(capacity int) LRUCache {
    cache := LRUCache{}
    cache.capacity = capacity
    cache.dict = make(map[int]*Node)
    cache.head = &Node{Key: 0, Value: 0}
    cache.tail = &Node{Key: 0, Value: 0}
    cache.head.Next = cache.tail
    cache.tail.Prev = cache.head
    cache.head.Prev = nil
    cache.tail.Next = cache.head
    cache.n = 0
    return cache
}

func (cache *LRUCache) deleteNode(n *Node) {
    n.Prev.Next = n.Next
    n.Next.Prev = n.Prev
}

func (cache *LRUCache) addHead(n *Node) {
    n.Next = cache.head.Next
    n.Next.Prev = n
    n.Prev = cache.head
    cache.head.Next = n
}

func (this *LRUCache) Get(key int) int {
    _, ok := this.dict[key]
    if ok {
        n := this.dict[key] 
        res := n.Value
        this.deleteNode(n)
        this.addHead(n)
        return res
    }
    return -1
}


func (this *LRUCache) Put(key int, value int)  {
    _, ok := this.dict[key]
    if ok {
        n := this.dict[key]
        n.Value = value
        this.deleteNode(n)
        this.addHead(n)
    } else {
        n := &Node{Key: key, Value: value}
        this.dict[key] = n
        if this.n < this.capacity {
            this.n += 1
            this.addHead(n)
        } else {
            delete(this.dict, this.tail.Prev.Key)
            this.deleteNode(this.tail.Prev)
            this.addHead(n)
        }
    }
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
