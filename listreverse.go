package piscine

func ListReverse(l *List) {
	if l.Head == nil || l.Head.Next == nil {
		return
	}

	var prev, next *NodeL
	current := l.Head

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.Head = prev
	l.Tail = l.Head
	for l.Tail.Next != nil {
		l.Tail = l.Tail.Next
	}
}
