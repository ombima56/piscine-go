package piscine

func ListClear(l *List) {
	if l.Head != nil {
		current := l.Head
		for current != nil {
			next := current.Next
			current = nil
			current = next
		}
	}
	l.Head = nil
	l.Tail = nil
}
