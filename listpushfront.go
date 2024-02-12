package piscine

type ExistingNodel = NodeL

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type LinkedList struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *LinkedList, data interface{}) {
	newNode := &NodeL{Data: data, Next: l.Head}
	l.Head = newNode
	if l.Tail == nil {
		l.Tail = newNode
	}
}
