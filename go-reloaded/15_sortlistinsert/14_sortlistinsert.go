package student

type NodeI struct {
	Data int
	Next *NodeI
}

func SortListInsert(l *NodeI, data_ref int) *NodeI{
	n := &NodeI{Data: data_ref}
	n.Next = nil

	if l == nil || data_ref <= l.Data {
		n.Next = l
		return n
	}

	temp := l

	for temp.Next != nil && temp.Next.Data < data_ref {
		temp = temp.Next
	}
	n.Next = temp.Next
	temp.Next = n
	return l
}
