package piscine

type nodeL struct {
	Data interface{}
	Next *nodeL
}

type List struct {
	Head *nodeL
	Tail *nodeL
}

func ListPushFront(l *List, data interface{}) {
	newNode := &nodeL{Data: data, Next: l.Head}
	l.Head = newNode
	if l.Tail == nil {
		l.Tail = newNode
	}
}
