package student

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {

	n := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}

	it := l.Head

	for it.Next != nil {
		it = it.Next
		if it.Next == nil {
			l.Tail = n
		}
	}
	it.Next = n
}

func ListRemoveIf(l *List, data_ref interface{}) {
	curr := l.Head
	var prev *NodeL = nil

	for curr != nil {

		if curr.Data != data_ref {
			prev = curr
			curr = curr.Next
			continue
		}

		curr = curr.Next
		if prev == nil {
			l.Head = curr
		} else {
			prev.Next = curr
		}
	}
}
