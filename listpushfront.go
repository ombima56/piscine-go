package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type list struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *list, data interface{}) {
	q := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = q
		return
	}
	q.Next = l.Head
	l.Head = q
}
