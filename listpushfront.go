package piscine

func ListPushFront(l *List, data interface{}) {
	q := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = q
		return
	}
	q.Next = l.Head
	l.Head = q
}
