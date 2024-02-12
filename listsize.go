package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
	if l.Head == nil {
		return 0
	}
	size := 1
	for l.Head.Next != nil {
		size++
		l.Head = l.Head.Next
	}
	return size
}

func ListPushFront(l *List, data interface{}) {
	q := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = q
		return
	}
	q.Next = l.Head
	l.Head = q
}
