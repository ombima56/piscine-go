package piscine

func listLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}
	nodes := l.Nodes()
	return nodes[len(nodes)-1].Data
}
