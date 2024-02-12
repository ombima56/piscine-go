package piscine

type Nodel struct {
	Data interface{}
	Next *Nodel
}

type List struct {
	Head *Nodel
	Tail *Nodel
}

func ListPushFront(l *List, data interface{}) {
	newNode := &Nodel{Data: data, Next: l.Head}
	l.Head = newNode
	if l.Tail == nil {
		l.Tail = newNode
	}
}
