package piscine

type ExistingNodel = nodeL

type nodeL struct {
	Data interface{}
	Next *nodeL
}

type LinkedList struct {
	Head *nodeL
	Tail *nodeL
}

func ListPushFront(l *LinkedList, data interface{}) {
	newNode := &nodeL{Data: data, Next: l.Head}
	l.Head = newNode
	if l.Tail == nil {
		l.Tail = newNode
	}
}
