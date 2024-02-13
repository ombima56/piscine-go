package piscine

type NodeK struct {
	Data interface{}
	Next *NodeK
}

type list struct {
	Head *NodeK
	Tail *NodeK
}

func ListLast(l *list) interface{} {
	if l.Head == nil {
		return nil
	}
	return l.Tail.Data
}
