type QueueNode struct {
    Value int
    Next *QueueNode
}

type Queue struct {
    Head *QueueNode
    Tail *QueueNode
}

func (queue *Queue) Push(value int) {
    if queue == nil {
        panic("Queue is nil")
    }
    newNode := &QueueNode{Value: value}
    if queue.Empty() {
        queue.Head = newNode
        queue.Tail = newNode
    } else {
        queue.Tail.Next = newNode
        queue.Tail = newNode
    }
}

func (queue *Queue) ExtractFront() (int, bool) {
    if queue == nil {
        panic("Queue is nil")
    }
    if queue.Empty() {
        return 0, false
    }
    front := queue.Head.Value
    if queue.Head == queue.Tail {
        queue.Head = nil
        queue.Tail = nil
    } else {
        queue.Head = queue.Head.Next
    }
    return front, true
}

func (queue *Queue) Empty() bool {
    if queue == nil {
        panic("Queue is nil")
    }
    if queue.Head == nil && queue.Tail == nil {
        return true
    }
    return false
}
