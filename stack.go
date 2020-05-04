package stack

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)

type Node struct {
    Value int
    Prev  *Node
}

type Stack struct {
    TopNode   *Node
}

func (s *Stack) Empty() bool {
    if s == nil {
        panic("nil stack")
    }
    return s.TopNode == nil
}

func (s *Stack) Top() int {
    if s == nil {
        panic("nil stack")
    }
    if s.TopNode == nil {
        panic("cannot get the top value of an empty stack")
    }
    return s.TopNode.Value
}

func (s *Stack) Push(value int) {
    if s == nil {
        panic("nil stack")
    }
    newNode := &Node{Value: value, Prev: s.TopNode}
    s.TopNode = newNode
}

func (s *Stack) Pop() int {
    if s == nil {
        panic("nil stack")
    }
    if s.TopNode == nil {
        panic("cannot pop the top value of an empty stack")
    }
    value := s.TopNode.Value
    s.TopNode = s.TopNode.Prev
    return value
}

type Queue struct {
    SPush  *Stack
    SPop   *Stack
}

func (q *Queue) Push(value int) {
    if q == nil {
        panic("nil queue")
    }
    if q.SPush == nil {
        q.SPush = &Stack{}
    }
    q.SPush.Push(value)
}

func (q *Queue) Pop() int {
    if q == nil {
        panic("nil queue")
    }
    if q.SPop.Empty() {
        q.swapStacks()
    }
    return q.SPop.Pop()
}

func (q *Queue) Front() int {
    if q == nil {
        panic("nil queue")
    }
    if q.SPop.Empty() {
        q.swapStacks()
    }
    return q.SPop.Top()
}

func (q *Queue) swapStacks() {
    for ; !q.SPush.Empty(); {
        v := q.SPush.Pop()
        q.SPop.Push(v)
    }
}
